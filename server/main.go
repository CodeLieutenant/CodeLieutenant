package main

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"os"
	"strconv"

	"github.com/malusev998/dusanmalusev/database"
	"github.com/malusev998/dusanmalusev/handlers"
	"github.com/malusev998/dusanmalusev/validator"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/gofiber/helmet"
)

func generateUniqueId() string {
	bytes := make([]byte, 64)

	read, err := rand.Read(bytes)

	if err != nil {
		return ""
	}

	if read != 64 {
		return ""
	}

	return base64.RawURLEncoding.EncodeToString(bytes)
}

func connectToDatabase() {
	err := database.Connect(os.Getenv("POSTGRES_DSN"))

	if err != nil {
		log.Fatal(err)
	}

	err = database.RunMigrations()

	if err != nil {
		log.Fatal(err)
	}
}

func setupFiberAppMiddleware(app *fiber.App) *fiber.App {
	app.Use(middleware.RequestID(middleware.RequestIDConfig{
		Generator: generateUniqueId,
	}))

	app.Use(middleware.Logger())

	app.Use(helmet.New())

	return app
}

func main() {

	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		log.Fatal(err)
	}

	validator.SetupValidator()

	connectToDatabase()

	app := setupFiberAppMiddleware(fiber.New(&fiber.Settings{
		ErrorHandler: validator.ErrorHandler,
	}))

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World 👋!")
	})

	api := app.Group("/api/v1")

	api.Use(cors.New(cors.Config{}))

	handlers.AddContactRoutes(api)

	log.Fatalf("Cannot start server on port %d: %v", port, app.Listen(port))
}
