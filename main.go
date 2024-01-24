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

	Queries = data.New(db)

	app := fiber.New()

	app.Static("/static", "./static")

	app.Get("/", handler.Index)
	app.Get("/signup", handler.SignUp)

	app.Get("/hello", handler.Hello)
	app.Get("/list", handler.List)

	log.Fatal(app.Listen("localhost:3000"))
}
