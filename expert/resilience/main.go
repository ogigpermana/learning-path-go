// File: expert/resilience/main.go
// Level: Expert
// Topik: Resilience Patterns (Circuit Breaker, Rate Limiter, Retry)
//
// Resilience = kemampuan sistem tetap berfungsi saat terjadi kegagalan.
//
// Patterns:
// 1. Retry: ulang operasi yang gagal
// 2. Circuit Breaker: stop request ke service yang down
// 3. Rate Limiter: batasi jumlah request
// 4. Timeout: batasi waktu tunggu
// 5. Bulkhead: isolasi resource (connection pool terpisah)

package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// ==========================================
// 1. RETRY PATTERN
// ==========================================

type RetryConfig struct {
	MaxAttempts int
	BaseDelay   time.Duration
	MaxDelay    time.Duration
}

func Retry(config RetryConfig, fn func() error) error {
	var lastErr error

	for attempt := 0; attempt < config.MaxAttempts; attempt++ {
		if attempt > 0 {
			delay := config.BaseDelay * (1 << attempt) // exponential backoff
			if delay > config.MaxDelay {
				delay = config.MaxDelay
			}
			fmt.Printf("  Retry %d/%d (waiting %v)...\n",
				attempt+1, config.MaxAttempts, delay)
			time.Sleep(delay)
		}

		err := fn()
		if err == nil {
			return nil // sukses
		}
		lastErr = err

		// Hentikan retry jika error tertentu
		if errors.Is(err, ErrNonRetryable) {
			fmt.Println("  Non-retryable error, stopping")
			break
		}
	}

	return fmt.Errorf("all %d attempts failed: %w", config.MaxAttempts, lastErr)
}

var ErrNonRetryable = errors.New("non-retryable error")

// ==========================================
// 2. CIRCUIT BREAKER
// ==========================================

type CircuitState int

const (
	StateClosed CircuitState = iota   // normal: request jalan
	StateOpen                         // rusak: request langsung ditolak
	StateHalfOpen                     // testing: coba 1 request
)

type CircuitBreaker struct {
	mu               sync.Mutex
	state            CircuitState
	failureCount     int
	threshold        int
	resetTimeout     time.Duration
	lastFailureTime  time.Time
	halfOpenSuccess  bool
}

func NewCircuitBreaker(threshold int, resetTimeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:       StateClosed,
		threshold:   threshold,
		resetTimeout: resetTimeout,
	}
}

func (cb *CircuitBreaker) Execute(fn func() error) error {
	cb.mu.Lock()

	// Cek state
	switch cb.state {
	case StateOpen:
		if time.Since(cb.lastFailureTime) > cb.resetTimeout {
			cb.state = StateHalfOpen
			fmt.Println("  [CB] State: OPEN -> HALF-OPEN")
		} else {
			cb.mu.Unlock()
			return ErrCircuitOpen
		}
	case StateHalfOpen:
		// Allow exactly 1 request
	}

	cb.mu.Unlock()

	// Execute
	err := fn()

	cb.mu.Lock()
	defer cb.mu.Unlock()

	if err != nil {
		cb.failureCount++
		cb.lastFailureTime = time.Now()

		if cb.failureCount >= cb.threshold {
			cb.state = StateOpen
			fmt.Printf("  [CB] State: CLOSED -> OPEN (failures: %d)\n", cb.failureCount)
		}
		return err
	}

	// Success
	if cb.state == StateHalfOpen {
		cb.state = StateClosed
		cb.failureCount = 0
		fmt.Println("  [CB] State: HALF-OPEN -> CLOSED (recovered)")
	}

	cb.failureCount = 0
	return nil
}

var ErrCircuitOpen = errors.New("circuit breaker is open")

// ==========================================
// 3. RATE LIMITER (Token Bucket)
// ==========================================

type RateLimiter struct {
	mu       sync.Mutex
	tokens   float64
	rate     float64  // tokens per second
	capacity float64  // max tokens
	lastRefill time.Time
}

func NewRateLimiter(rate, capacity float64) *RateLimiter {
	return &RateLimiter{
		tokens:   capacity,
		rate:     rate,
		capacity: capacity,
		lastRefill: time.Now(),
	}
}

func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	// Refill tokens
	now := time.Now()
	elapsed := now.Sub(rl.lastRefill).Seconds()
	rl.tokens += elapsed * rl.rate
	if rl.tokens > rl.capacity {
		rl.tokens = rl.capacity
	}
	rl.lastRefill = now

	// Check if allowed
	if rl.tokens >= 1 {
		rl.tokens--
		return true
	}
	return false
}

// ==========================================
// DEMO
// ==========================================

func main() {
	fmt.Println("=== 1. RETRY PATTERN ===")
	attempts := 0
	err := Retry(RetryConfig{MaxAttempts: 3, BaseDelay: 50 * time.Millisecond, MaxDelay: 200 * time.Millisecond}, func() error {
		attempts++
		fmt.Printf("  Attempt %d...\n", attempts)
		if attempts < 3 {
			return errors.New("temporary error")
		}
		return nil
	})
	fmt.Printf("  Result: %v\n\n", err)

	fmt.Println("=== 2. CIRCUIT BREAKER ===")
	cb := NewCircuitBreaker(2, 100*time.Millisecond)

	// Simulasi error
	for i := 0; i < 3; i++ {
		err := cb.Execute(func() error {
			return errors.New("service error")
		})
		fmt.Printf("  Request %d: %v\n", i+1, err)
	}

	// Request ketika circuit open
	err = cb.Execute(func() error {
		return nil
	})
	fmt.Printf("  Request saat OPEN: %v\n", err)

	// Tunggu reset
	fmt.Println("  Menunggu reset...")
	time.Sleep(150 * time.Millisecond)

	// Half-open: sukses -> close
	err = cb.Execute(func() error {
		return nil
	})
	fmt.Printf("  Request setelah reset: %v\n", err)
	fmt.Println()

	fmt.Println("=== 3. RATE LIMITER ===")
	rl := NewRateLimiter(5, 5) // 5 request per detik

	for i := 0; i < 10; i++ {
		allowed := rl.Allow()
		fmt.Printf("  Request %d: allowed=%v\n", i+1, allowed)
		time.Sleep(100 * time.Millisecond) // 100ms = 10/sec, tapi rate=5/sec
	}
	fmt.Println()

	fmt.Println("=== Best Practices ===")
	fmt.Println("1. Retry dengan exponential backoff + jitter")
	fmt.Println("2. Circuit breaker untuk service external")
	fmt.Println("3. Rate limiter untuk API protection")
	fmt.Println("4. Timeout untuk semua external calls")
	fmt.Println("5. Bulkhead: pool terpisah per service")
	fmt.Println()
	fmt.Println("Library production:")
	fmt.Println("- github.com/sony/gobreaker (circuit breaker)")
	fmt.Println("- github.com/ulule/limiter (rate limiter)")
	fmt.Println("- github.com/avast/retry-go (retry)")
	fmt.Println("- github.com/cenkalti/backoff (backoff)")
}

/*
Demo retry dengan jitter (random delay):
func jitter(delay time.Duration) time.Duration {
    return time.Duration(float64(delay) * (0.5 + rand.Float64()))
}
*/