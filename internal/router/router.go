package router

import (
	b "fiber/internal/booking"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/sync/errgroup"
	"os"
)

type BookingBody struct {
	ClientID   primitive.ObjectID `json:"clientId,omitempty" bson:"clientId,omitempty" validate:"required"`
	TimeSlotID primitive.ObjectID `json:"timeSlotId,omitempty" bson:"timeSlotId,omitempty" validate:"required"`
}

func (r *router) Start(port int) {
	r.app.Use(limiter.New())
	r.app.Use(logger.New())

	// JWT Middleware
	r.app.Use(jwtware.New(jwtware.Config{
		SigningKey:               []byte(os.Getenv("ACCESS_TOKEN_SECRET")),
		SigningMethod:            "HS256",
		ContextKey:               "data",
		TokenLookup:              "header:Authorization",
		AuthScheme:               "Bearer",
	}))

	r.app.Get("/monitor", monitor.New())

	r.app.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).
			JSON(fiber.Map{"ok": "true", "msg": "ping", "data": "{}"})
	})

	api := r.app.Group("/api", func(c *fiber.Ctx) error {
		data := c.Locals("data").(*jwt.Token)
		claims := data.Claims.(jwt.MapClaims)
	 	user := claims["data"].(map[string]interface{})
		id, err := primitive.ObjectIDFromHex(user["_id"].(string))
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).
				JSON(fiber.Map{"ok": "false", "msg": err.Error()})
		}
		count, err := r.mongo.CountUser(c.Context(), &b.User{ID: id})
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).
				JSON(fiber.Map{"ok": "false", "msg": err.Error()})
		}
		if count == 0 {
			return  c.Status(fiber.StatusUnauthorized).
				JSON(fiber.Map{"ok": "false", "msg": err.Error()})
		}
	 	c.Next()
		return nil
	})

	api.Post("/booking", func(c *fiber.Ctx) error {
		booking := new(BookingBody)
		if err := c.BodyParser(booking); err != nil {
			return  c.Status(fiber.StatusBadRequest).
				JSON(fiber.Map{"ok": "false", "msg": err.Error()})
		}


		validate := validator.New()
		err := validate.Struct(booking)
		if err != nil {
			return  c.Status(fiber.StatusBadRequest).
				JSON(fiber.Map{"ok": "false", "msg": err.Error()})
		}

		g, gctx := errgroup.WithContext(c.Context())
		var result *b.Booking
		g.Go(func() error {
			data, err := r.mongo.CreateBooking(gctx, &b.Booking{
				ClientID:   booking.ClientID,
				TimeSlotID: booking.TimeSlotID,
			})
			if err != nil {
				return err
			}
			result = data
			return nil
		})
		if err := g.Wait(); err != nil {
			return  c.Status(fiber.StatusInternalServerError).
				JSON(fiber.Map{"ok": "false", "msg": err.Error()})
		}
		return c.Status(fiber.StatusCreated).
			JSON(fiber.Map{"ok": "true", "msg": "", "data": ""})
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
