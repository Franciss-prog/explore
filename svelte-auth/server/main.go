package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app_2 := fiber.New()
	// use cors here
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	app.Post("/auth", func(c *fiber.Ctx) error {
		Username := c.FormValue("username")
		Password := c.FormValue("password")
		// form validation again
		if Username == "" || Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Username and Password are required",
				"status":  401,
			})
		}
		return c.JSON(fiber.Map{
			"data":    fiber.Map{"username": Username, "password": Password},
			"message": "Login Success",
			"status":  200,
		})
	})
	app_2.Use(cors.New(cors.Config{}))
	app_2.Get("/", func(c *fiber.Ctx) error {
		response := c.SendString("Hello, World!")
		return response
	})

	go func() {
		if err := app.Listen(":3000"); err != nil {
			panic(err)

		}
	}()
	go func() {
		if err := app_2.Listen(":3001"); err != nil {
			panic(err)

		}
	}()
	select {}
}
