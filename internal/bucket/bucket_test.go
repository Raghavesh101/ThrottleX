package bucket

import (
	"testing"
	"time"
)

func TestAllowBurstThenReject(t *testing.T) {
	b := NewBucket(1, 10) //10 tokens per second

	for i := 0; i < 5; i++ {
		if !b.Allow() {
			t.Fatalf("request %d: expected allow, got reject", i)
		}
	}
	if !b.Allow() {
		t.Fatalf("request 6: expected reject, got allow")
	}
}

func TestAllowRefillAfterWait(t *testing.T) {
	b := NewBucket(1, 10) //10 tokens per second

	if !b.Allow() {
		t.Fatalf("request 1: expected allow, got reject")
	}
	if !b.Allow() {
		t.Fatalf("request 2: expected reject(bucket empty), got allow")
	}
	time.Sleep(250 * time.Millisecond) // wait for 0.25 seconds, should refill 2 tokens

	if !b.Allow() {
		t.Fatalf("request 3: expected allow after refill, got reject")
	}
}
