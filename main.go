package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mohamedimrane/twotor/handler"
)

func main() {
	app := fiber.New()

	app.Static("/static", "./static")

	app.Get("/", handler.Hello)
	app.Get("/list", handler.List)

	log.Fatal(app.Listen("localhost:3000"))
}
