package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/khajamoddin/collections/collections"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	const key = "active-sessions:region:eu-west-1"

	// 1) Load current active session IDs from Redis SET
	rawIDs, err := rdb.SMembers(ctx, key).Result()
	if err != nil {
		// Just log error for demo purposes if Redis isn't running
		log.Printf("redis SMEMBERS failed (is redis running?): %v", err)
		// return // don't return so we can demo the collections logic below even without redis
	}

	active := collections.NewSetFromSlice(rawIDs)

	log.Printf("Loaded %d active sessions from Redis", active.Len())

	// 2) Simulate a few new sessions we observed in this process
	newSessions := []string{"sess-101", "sess-102", "sess-103"}
	for _, id := range newSessions {
		if active.Add(id) {
			if err := rdb.SAdd(ctx, key, id).Err(); err != nil {
				log.Printf("redis SADD failed: %v", err)
			}
			log.Printf("Added new session %s", id)
		}
	}

	// 3) Simulate sessions that timed out locally and need removal
	expiredLocally := []string{"sess-000", "sess-101"}
	expiredSet := collections.NewSet[string]()
	for _, s := range expiredLocally {
		expiredSet.Add(s)
	}

	// Sessions to actually remove from Redis = active ∩ expired
	toRemove := active.Intersection(expiredSet)
	if toRemove.Len() > 0 {
		var args []any
		for id := range toRemove.All() {
			args = append(args, id)
		}
		if err := rdb.SRem(ctx, key, args...).Err(); err != nil {
			log.Printf("redis SREM failed: %v", err)
		}
		log.Printf("Removed %d expired sessions", toRemove.Len())
	}

	// 4) Show all active sessions via iterator
	log.Println("Current active sessions:")
	for id := range active.All() {
		fmt.Println(" -", id)
	}

	// Keep process alive briefly to inspect logs
	time.Sleep(100 * time.Millisecond)
}
