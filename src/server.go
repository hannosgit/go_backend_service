package main

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Post("/auth/authenticate", func(c *fiber.Ctx) error {
		time.Sleep(250 * time.Millisecond)
		c.Set("Content-type", "application/json; charset=utf-8")
		return c.SendString("{\"token\": \"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c\"}")
	})
	app.Get("/delivery/:deliveryId", func(c *fiber.Ctx) error {
		time.Sleep(250 * time.Millisecond)
		c.Set("Content-type", "application/json; charset=utf-8")
		return c.SendString("{\"deliveryId\": \"28461\", \"restaurantName\": \"Ristorante Luigi\",\"deliveryDate\": \"2023-03-04\", \"totalItems\": 5}")
	})
	app.Get("/bill/:billId", func(c *fiber.Ctx) error {
		time.Sleep(250 * time.Millisecond)
		c.Set("Content-type", "application/json; charset=utf-8")
		return c.SendString("{\"billId\": 123, \"totalPrice\": 1250, \"accountId\": 1250, \"firstName\": \"John\", \"lastName\": \"Doe\", \"address\": \"Fake Street\"}")
	})

	app.Listen(":3000")
}
