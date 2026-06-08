package bucket

import (
	"time"
)

type Bucket struct {
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
	now := time.Now()
	elapsed := now.Sub(b.lastRefill).Seconds()
	b.tokens = min(b.tokens+elapsed*b.refillRate, b.capacity)
	if b.tokens >= 1 {
		b.tokens--
		return true
	}
	return false
}
