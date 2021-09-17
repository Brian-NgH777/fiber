package router

import (
	b "fiber/internal/booking"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"golang.org/x/sync/errgroup"
)

func (r *router) Start(port int) {

	r.app.Use(cors.New())
	r.app.Use(limiter.New())
	r.app.Use(logger.New())

	r.app.Get("/monitor", monitor.New())

	r.app.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(201).JSON("{}")
	})

	api := r.app.Group("/api", MiddlewareAuth)

	api.Post("/api/booking", func(c *fiber.Ctx) error {
		booking := new(b.Booking)
		if err := c.BodyParser(booking); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		g, gctx := errgroup.WithContext(c.Context())
		var result *b.Booking
		g.Go(func() error {
			data, err := r.mongo.CreateBooking(gctx, booking)
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

	r.app.Listen(fmt.Sprintf(":%d", port))
}

func New() *router {
	app := fiber.New()
	mongo := b.Connect()
	return &router{app: app, mongo: mongo}
}

type router struct {
	app            *fiber.App
	mongo          *b.MongoInstance
}
