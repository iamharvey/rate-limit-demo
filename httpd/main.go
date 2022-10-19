package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"log"
	"net/http"
)

func main() {
	// 100/s, burst of 200.
	limiter := rate.NewLimiter(rate.Limit(100), 200)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "too many requests", http.StatusTooManyRequests)
			return
		}

		fmt.Fprintf(w, "Hi, Gophers!")
	})

	log.Println("starting server at port 4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
