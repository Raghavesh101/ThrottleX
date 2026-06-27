package bucket

import (
	"sync"
	"time"
)

type Bucket struct {
	mu         sync.Mutex
	capacity   float64
	refillRate float64
	tokens     float64
	lastRefill time.Time
}

func NewBucket(capacity, refillRate float64) *Bucket {
	return &Bucket{
		capacity:   capacity,
		refillRate: refillRate,
		tokens:     capacity,
		lastRefill: time.Now(),
	}
}

func (b *Bucket) Allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(b.lastRefill).Seconds()
	b.tokens = min(b.tokens+elapsed*b.refillRate, b.capacity)
	b.lastRefill = now
	if b.tokens >= 1 {
		b.tokens--
		return true
	}
	return false
}
