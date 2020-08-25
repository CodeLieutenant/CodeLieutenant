package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World 👋!")
	})

	app.Listen(port)
}
