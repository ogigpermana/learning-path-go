// File: expert/pubsub/main.go
// Level: Expert
// Topik: Pub/Sub Pattern (Publish-Subscribe)
//
// Pub/Sub adalah pola komunikasi asynchronous:
// - Publisher: mengirim pesan ke channel/topic
// - Subscriber: menerima pesan dari channel/topic
// - Decoupled: publisher tidak tahu subscriber (dan sebaliknya)
//
// Implementasi: in-memory (contoh), RabbitMQ, Kafka, Redis Pub/Sub, NATS

package main

import (
	"fmt"
	"sync"
	"time"
)

// ==========================================
// IN-MEMORY PUB/SUB (sederhana)
// ==========================================

// Message adalah pesan yang dikirim
type Message struct {
	Topic     string
	Payload   interface{}
	Timestamp time.Time
}

// Subscriber adalah fungsi yang menangani pesan
type Subscriber func(msg Message)

// PubSub adalah message broker sederhana
type PubSub struct {
	subscribers map[string][]Subscriber
	mu          sync.RWMutex
}

func NewPubSub() *PubSub {
	return &PubSub{
		subscribers: make(map[string][]Subscriber),
	}
}

// Subscribe: daftarkan subscriber untuk topic tertentu
func (ps *PubSub) Subscribe(topic string, fn Subscriber) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ps.subscribers[topic] = append(ps.subscribers[topic], fn)
	fmt.Printf("[PubSub] Subscriber registered for topic: %s\n", topic)
}

// Publish: kirim pesan ke semua subscriber topic
func (ps *PubSub) Publish(topic string, payload interface{}) {
	ps.mu.RLock()
	subs := ps.subscribers[topic]
	ps.mu.RUnlock()

	msg := Message{
		Topic:     topic,
		Payload:   payload,
		Timestamp: time.Now(),
	}

	fmt.Printf("[PubSub] Publishing to %s: %v\n", topic, payload)

	// Kirim ke semua subscriber (async)
	for _, sub := range subs {
		go sub(msg)
	}
}

// Close: cleanup
func (ps *PubSub) Close() {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ps.subscribers = make(map[string][]Subscriber)
	fmt.Println("[PubSub] Closed all subscribers")
}

// ==========================================
// DEMO
// ==========================================

func main() {
	ps := NewPubSub()
	defer ps.Close()

	// 1. SUBSCRIBE - daftarkan handlers
	fmt.Println("=== Subscribe ===")

	// Email notification handler
	ps.Subscribe("user.registered", func(msg Message) {
		user := msg.Payload.(string)
		fmt.Printf("[Email] Sending welcome email to %s...\n", user)
		time.Sleep(50 * time.Millisecond)
		fmt.Printf("[Email] Welcome email sent to %s\n", user)
	})

	// Logging handler
	ps.Subscribe("user.registered", func(msg Message) {
		user := msg.Payload.(string)
		fmt.Printf("[Log] User registered at %s: %s\n",
			msg.Timestamp.Format(time.RFC3339), user)
	})

	// Dashboard analytics handler
	ps.Subscribe("user.registered", func(msg Message) {
		fmt.Printf("[Analytics] Incrementing registration counter...\n")
	})

	// Order handler
	ps.Subscribe("order.created", func(msg Message) {
		orderID := msg.Payload.(int)
		fmt.Printf("[Order] Processing order #%d...\n", orderID)
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("[Order] Order #%d confirmed\n", orderID)
	})

	fmt.Println()

	// 2. PUBLISH - kirim pesan
	fmt.Println("=== Publish Events ===")
	ps.Publish("user.registered", "Anggi")
	time.Sleep(100 * time.Millisecond)

	ps.Publish("user.registered", "Budi")
	time.Sleep(100 * time.Millisecond)

	ps.Publish("order.created", 1001)
	time.Sleep(150 * time.Millisecond)

	fmt.Println()
	fmt.Println("=== Use Cases ===")
	fmt.Println("1. Event-driven architecture")
	fmt.Println("2. Microservices communication")
	fmt.Println("3. Real-time notifications")
	fmt.Println("4. Log aggregation")
	fmt.Println("5. CQRS (Command Query Responsibility Segregation)")
}

/*
Production Pub/Sub:
- RabbitMQ: go get github.com/rabbitmq/amqp091-go
- Kafka: go get github.com/segmentio/kafka-go
- Redis Pub/Sub: go get github.com/go-redis/redis/v8
- NATS: go get github.com/nats-io/nats.go
- Google Pub/Sub: cloud.google.com/go/pubsub
*/