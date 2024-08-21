package main

import (
	"fmt"
	"go-tutorial/db"
	"go-tutorial/user"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		log.Fatalf("cannot connected to db %v", err)
	}

	fmt.Println(database.Config)

	repo := user.NewRepository(database)
	err = repo.Migration()
	if err != nil {
		log.Fatal(err)
	}
	service := user.NewService(repo)
	handler := user.NewHandler(service)

	app := fiber.New()
	app.Get("/users/:id", handler.Get)
	app.Post("/users", handler.Create)

	app.Listen(":8000")
}
