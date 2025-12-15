---
layout: default
title: Recipes
nav_order: 6
---

# Recipes

Real-world patterns using the `collections` library.

## 1. Unique Item Processing (Set)

Deduplicate a stream of IDs (e.g., from a CSV or API) before processing.

```go
import "github.com/khajamoddin/collections/collections"

func ProcessUniqueIDs(ids []string) {
    seen := collections.NewSet[string]()
    for _, id := range ids {
        if !seen.Has(id) {
            seen.Add(id)
            // process id...
        }
    }
}
```

## 2. LRU Cache Building Block (OrderedMap)

Build a basic Least-Recently-Used (LRU) cache using `OrderedMap`'s stability.

```go
// Note: Real LRU needs move-to-front logic. OrderedMap preserves insertion order.
// To make it LRU, you delete and re-insert on access.

func Access(om *collections.OrderedMap[string, Data], key string) Data {
    if v, ok := om.Get(key); ok {
        om.Delete(key)
        om.Set(key, v) // Move to back (most recently used)
        return v
    }
    return Data{}
}
```

## 3. Task Scheduler (PriorityQueue)

Manage tasks with different priorities.

```go
type Task struct {
    Priority int
    Name     string
}

func Schedule() {
    // Min-heap: Lower priority number processed first
    pq := collections.NewPriorityQueue[Task](func(a, b Task) bool {
        return a.Priority < b.Priority
    })
    
    pq.Push(Task{Priority: 10, Name: "Low"})
    pq.Push(Task{Priority: 1, Name: "Critical"})
    
    // Process "Critical" first
    if task, ok := pq.Pop(); ok {
        fmt.Println("Processing:", task.Name)
    }
}
```

## 4. Multi-Index (MultiMap)

Index objects by a non-unique field (e.g., users by role).

```go
func IndexUsersByRole(users []User) *collections.MultiMap[string, User] {
    mm := collections.NewMultiMap[string, User]()
    for _, u := range users {
        mm.Add(u.Role, u)
    }
    return mm
}
// Get all "admins":
// admins := mm.Get("admin") // returns []User
```
