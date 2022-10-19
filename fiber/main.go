package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"log"
	"time"
)

func main() {
	app := fiber.New()

	app.Use(limiter.New(limiter.Config{
		// Max number of recent connections during `Expiration` seconds before sending a 429 response
		Max: 20,

		// Expiration is the time on how long to keep records of requests in memory.
		Expiration: 30 * time.Second,

		// LimitReached is called when a request hits the limit.
		LimitReached: func(c *fiber.Ctx) error {
			fmt.Println("too many requests")
			return c.SendStatus(fiber.StatusTooManyRequests)
		},

		// When set to true, requests with StatusCode >= 400 won't be counted.
		SkipFailedRequests: false,

		// When set to true, requests with StatusCode < 400 won't be counted.
		SkipSuccessfulRequests: false,

		// LimiterMiddleware is the struct that implements a limiter middleware.
		// limiter.SlidingWindow{} will take into account the previous window(if there was any)
		LimiterMiddleware: limiter.FixedWindow{},
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi, Gophers!")
	})

	log.Fatal(app.Listen(":3000"))
}
