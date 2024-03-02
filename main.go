package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/smallbatch-apps/trivialize-go/database"
	"github.com/smallbatch-apps/trivialize-go/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to the API!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api", welcome)
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
}

func main() {
	database.Connect()
	app := fiber.New()
	setupRoutes(app)
	godotenv.Load()
	log.Fatal(app.Listen(":3000"))
}
