package main

import (
	"log"
	"net/http"

	"github.com/Raghavesh101/throttlex/internal/middleware"
)

func main() {
	// Create a new rate limiter with a limit of 5 requests per second and a burst size of 10
	l := limiter.NewLimiter(5, 1)
	mux := http.NewServeMux()
	// Wrap the handler with the rate limiting middleware
	mux.HandleFunc("/", (func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok\n"))
	}))
	handler := middleware.RateLimit(l, mux)
	addr := ":8080"
	log.Printf("Throttlx server on %s", addr)
	log.Fatal(http.ListenAndServe(addr, handler))

}
