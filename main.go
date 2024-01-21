package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mohamedimrane/twotor/handlers"
)

func main() {
	app := fiber.New()

	app.Static("/static", "./static")

	app.Get("/", handlers.Hello)
	app.Get("/list", handlers.List)

	log.Fatal(app.Listen("localhost:3000"))
}
