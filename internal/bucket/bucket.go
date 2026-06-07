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
