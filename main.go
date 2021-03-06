package main

import (
	b "fiber/booking"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"golang.org/x/sync/errgroup"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	// Connect to the database
	if err := b.Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(201).JSON("do re mi fa son")
	})

	app.Post("/api/booking", func(c *fiber.Ctx) error {
		booking := new(b.Booking)
		if err := c.BodyParser(booking); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		g, gctx := errgroup.WithContext(c.Context())
		var result *b.Booking
		g.Go(func() error {
			data, err := b.CreateBooking(gctx, booking)
			if err != nil {
				return err
			}
			result = data
			return nil

		})
		if err := g.Wait(); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.Status(201).JSON(result)
	})

	app.Listen(":" + port)
}
