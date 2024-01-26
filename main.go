package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"

	"github.com/gofiber/fiber/v2"
	"github.com/mohamedimrane/twotor/data"
	"github.com/mohamedimrane/twotor/handler"
	"github.com/mohamedimrane/twotor/middleware"
)

var Queries *data.Queries

func main() {
	// Load environment varibles from file
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	// Database setup
	db, err := sql.Open("sqlite3", "database.sqlite")
	if err != nil {
		log.Fatalln(err)
	}
	queries := data.New(db)

	// Server setup
	app := fiber.New()

	app.Static("/static", "./static")

	hw := handler.New(queries)
	mw := middleware.New(queries)

	// Unauthenticated routes
	app.Get("/", mw.Unauthenticated, hw.Index)
	app.Get("/signup", mw.Unauthenticated, hw.SignUp)
	app.Post("/api/create-user", mw.Unauthenticated, hw.CreateUser)

	// Autenticated routes
	app.Get("/home", mw.Authenticated, hw.Home)

	app.Get("/hello", handler.Hello)
	app.Get("/list", handler.List)

	port := os.Getenv("PORT")
	log.Fatal(app.Listen("localhost:" + port))
}
