package main

import (
	"context"
	"log"
	"time"

	"github.com/khajamoddin/collections/collections"
	"github.com/segmentio/kafka-go"
)

type retryMessage struct {
	Msg         kafka.Message
	NextAttempt time.Time
	Attempt     int
}

func main() {
	ctx := context.Background()

	// Kafka reader (adjust brokers and topic to your env)
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "input-events",
		GroupID: "collections-retry-example",
	})
	defer reader.Close()

	// Simple "ready to process now" queue
	ready := collections.NewDeque[retryMessage]()

	// Priority queue for scheduled retries ordered by NextAttempt
	retries := collections.NewPriorityQueue[retryMessage](
		func(a, b retryMessage) bool {
			return a.NextAttempt.Before(b.NextAttempt)
		},
	)

	log.Println("Starting kafka-retry-queue example...")

	// Use a shorter ticker and deadline for demo responsiveness
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	// Stop after a few seconds for demo purposes
	timeout := time.After(5 * time.Second)

	for {
		select {
		case <-timeout:
			log.Println("Demo time limit reached")
			return
		case <-ctx.Done():
			log.Println("Context cancelled, exiting...")
			return

		case <-ticker.C:
			// 1) Move due retries into ready queue
			now := time.Now()
			for retries.Len() > 0 {
				top, ok := retries.Peek()
				if !ok || top.NextAttempt.After(now) {
					break
				}
				msg, _ := retries.Pop()
				ready.PushBack(msg)
			}

			// 2) Consume new messages (simulated here if no Kafka)
			_ = reader.SetReadDeadline(time.Now().Add(10 * time.Millisecond))
			msg, err := reader.ReadMessage(ctx)
			if err == nil {
				ready.PushBack(retryMessage{
					Msg:         msg,
					NextAttempt: time.Now(),
					Attempt:     0,
				})
			} else {
				// Simulate sporadic messages for demo if Kafka is down
				if time.Now().UnixNano()%10 == 0 {
					ready.PushBack(retryMessage{
						Msg:         kafka.Message{Offset: int64(time.Now().Unix())},
						NextAttempt: time.Now(),
						Attempt:     0,
					})
				}
			}

			// 3) Process ready messages
			for ready.Len() > 0 {
				rm, ok := ready.PopFront()
				if !ok {
					break
				}
				if processMessage(rm.Msg) {
					log.Printf("Processed message at offset %d successfully", rm.Msg.Offset)
					continue
				}

				// Failed: schedule retry with exponential backoff
				rm.Attempt++
				backoff := time.Duration(1<<rm.Attempt) * time.Millisecond * 100 // fast backoff for demo
				if backoff > 5*time.Second {
					backoff = 5 * time.Second
				}
				rm.NextAttempt = time.Now().Add(backoff)
				retries.Push(rm)

				log.Printf("Scheduled retry #%d for offset %d in %s",
					rm.Attempt, rm.Msg.Offset, backoff)
			}
		}
	}
}

// processMessage simulates a flaky handler
func processMessage(msg kafka.Message) bool {
	// fail some messages
	return msg.Offset%3 != 0
}
