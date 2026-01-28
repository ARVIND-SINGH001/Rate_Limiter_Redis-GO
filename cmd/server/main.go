package main

import (
	"log"
	"net/http"

	"rate-limiter/internal/middleware"
	"rate-limiter/internal/redis"
)

func main() {
	// 1. Initialize Redis Cloud connection
	redis.Init()

	// 2. Create HTTP router
	mux := http.NewServeMux()

	// 3. Register rate-limited API route
	mux.Handle(
		"/api",
		middleware.RateLimiter(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Request allowed\n"))
			}),
		),
	)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(
		"Distributed Rate Limiter API\n\n" +
		"GET /api  -> rate-limited endpoint\n\n" +
		"After limit -> returns HTTP 429\n",
	))
})


	// 4. Start HTTP server
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
