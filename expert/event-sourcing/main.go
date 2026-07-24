// File: expert/event-sourcing/main.go
// Level: Expert
// Topik: Event Sourcing
//
// Event Sourcing: menyimpan state sebagai urutan events (bukan state terkini).
// Keuntungan: audit trail lengkap, time travel, replay events.
//
// Konsep:
// - Event: sesuatu yang terjadi di masa lalu (immutable)
// - Event Store: database untuk menyimpan events
// - Aggregate: kumpulan events untuk satu entity
// - Projection: membaca events untuk menghasilkan state terkini

package main

import (
	"fmt"
	"time"
)

// ==========================================
// EVENT DEFINITIONS
// ==========================================

// Event adalah base interface untuk semua event
type Event interface {
	EventType() string
	Timestamp() time.Time
}

// Events spesifik untuk Bank Account
type AccountOpened struct {
	At time.Time
}

func (e AccountOpened) EventType() string { return "AccountOpened" }
func (e AccountOpened) Timestamp() time.Time { return e.At }

type MoneyDeposited struct {
	Amount float64
	At     time.Time
}

func (e MoneyDeposited) EventType() string { return "MoneyDeposited" }
func (e MoneyDeposited) Timestamp() time.Time { return e.At }

type MoneyWithdrawn struct {
	Amount float64
	At     time.Time
}

func (e MoneyWithdrawn) EventType() string { return "MoneyWithdrawn" }
func (e MoneyWithdrawn) Timestamp() time.Time { return e.At }

type AccountFrozen struct {
	Reason string
	At     time.Time
}

func (e AccountFrozen) EventType() string { return "AccountFrozen" }
func (e AccountFrozen) Timestamp() time.Time { return e.At }

// ==========================================
// AGGREGATE
// ==========================================

// BankAccount adalah aggregate: merekonstruksi state dari events
type BankAccount struct {
	ID      string
	Balance float64
	IsOpen  bool
	IsFrozen bool
	Events  []Event // event history
}

// ApplyEvent: mengupdate state berdasarkan event
// Ini adalah inti dari Event Sourcing: state = fold(events)
func (a *BankAccount) ApplyEvent(event Event) {
	switch e := event.(type) {
	case AccountOpened:
		a.IsOpen = true
	case MoneyDeposited:
		a.Balance += e.Amount
	case MoneyWithdrawn:
		a.Balance -= e.Amount
	case AccountFrozen:
		a.IsFrozen = true
	}
}

// LoadFromHistory: rekonstruksi state dari semua events
func (a *BankAccount) LoadFromHistory(events []Event) {
	for _, event := range events {
		a.ApplyEvent(event)
		a.Events = append(a.Events, event)
	}
}

// ==========================================
// COMMANDS
// ==========================================

// OpenAccount: command untuk membuka account
func (a *BankAccount) OpenAccount(id string) {
	event := AccountOpened{At: time.Now()}
	a.ApplyEvent(event)
	a.Events = append(a.Events, event)
	a.ID = id
	fmt.Printf("[CMD] Account %s opened\n", id)
}

// Deposit: command untuk deposit uang
func (a *BankAccount) Deposit(amount float64) error {
	if !a.IsOpen {
		return fmt.Errorf("account is closed")
	}
	if a.IsFrozen {
		return fmt.Errorf("account is frozen")
	}
	if amount <= 0 {
		return fmt.Errorf("amount must be positive")
	}

	event := MoneyDeposited{Amount: amount, At: time.Now()}
	a.ApplyEvent(event)
	a.Events = append(a.Events, event)
	fmt.Printf("[CMD] Deposited %.0f, balance: %.0f\n", amount, a.Balance)
	return nil
}

// Withdraw: command untuk tarik uang
func (a *BankAccount) Withdraw(amount float64) error {
	if !a.IsOpen {
		return fmt.Errorf("account is closed")
	}
	if a.IsFrozen {
		return fmt.Errorf("account is frozen")
	}
	if amount <= 0 {
		return fmt.Errorf("amount must be positive")
	}
	if amount > a.Balance {
		return fmt.Errorf("insufficient balance: %.0f < %.0f", a.Balance, amount)
	}

	event := MoneyWithdrawn{Amount: amount, At: time.Now()}
	a.ApplyEvent(event)
	a.Events = append(a.Events, event)
	fmt.Printf("[CMD] Withdrew %.0f, balance: %.0f\n", amount, a.Balance)
	return nil
}

// Freeze: command untuk freeze account
func (a *BankAccount) Freeze(reason string) {
	event := AccountFrozen{Reason: reason, At: time.Now()}
	a.ApplyEvent(event)
	a.Events = append(a.Events, event)
	fmt.Printf("[CMD] Account frozen: %s\n", reason)
}

// ==========================================
// MAIN
// ==========================================

func main() {
	fmt.Println("=== Event Sourcing Demo ===")
	fmt.Println()

	// 1. Buka account
	account := &BankAccount{}
	account.OpenAccount("ACC-001")

	// 2. Lakukan transaksi
	account.Deposit(1000)
	account.Deposit(500)
	account.Withdraw(200)
	account.Withdraw(100)

	// 3. Tampilkan state terkini
	fmt.Println("\n=== Current State ===")
	fmt.Printf("Account: %s\n", account.ID)
	fmt.Printf("Balance: %.0f\n", account.Balance)
	fmt.Printf("Open: %v, Frozen: %v\n", account.IsOpen, account.IsFrozen)
	fmt.Printf("Total events: %d\n", len(account.Events))

	// 4. Tampilkan event history (audit trail)
	fmt.Println("\n=== Event History (Audit Trail) ===")
	for i, event := range account.Events {
		fmt.Printf("  %d. %s at %s\n",
			i+1,
			event.EventType(),
			event.Timestamp().Format("15:04:05.000"))
	}

	// 5. TIME TRAVEL: rekonstruksi state pada titik tertentu
	fmt.Println("\n=== Time Travel ===")
	historical := &BankAccount{}
	// "Putar ulang" hanya 3 event pertama
	historical.LoadFromHistory(account.Events[:3])
	fmt.Printf("State after 3 events: balance=%.0f\n", historical.Balance)

	// 6. Freeze account
	fmt.Println()
	account.Freeze("Suspicious activity")
	account.Deposit(100) // akan error karena frozen

	fmt.Println()
	fmt.Println("=== Keuntungan Event Sourcing ===")
	fmt.Println("1. Audit trail lengkap")
	fmt.Println("2. Time travel / replay events")
	fmt.Println("3. Debugging lebih mudah")
	fmt.Println("4. CQRS (command/query separation)")
	fmt.Println("5. Event-driven microservices")
}

/*
Production Event Store:
- EventStoreDB (geteventstore.com)
- Kafka sebagai event store
- PostgreSQL dengan event table
- DynamoDB / CosmosDB
*/