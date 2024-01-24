package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gofiber/fiber/v2"
	"github.com/mohamedimrane/twotor/data"
	"github.com/mohamedimrane/twotor/handler"
)

var Queries *data.Queries

func main() {
	db, err := sql.Open("sqlite3", "database.sqlite")
	if err != nil {
		log.Fatalln(err)
	}
	queries := data.New(db)

	app := fiber.New()

	app.Static("/static", "./static")

	hw := handler.New(queries)

	app.Get("/", hw.Index)
	app.Get("/signup", hw.SignUp)

	app.Post("/api/create-user", hw.CreateUser)

	app.Get("/hello", handler.Hello)
	app.Get("/list", handler.List)

	log.Fatal(app.Listen("localhost:3000"))
}
