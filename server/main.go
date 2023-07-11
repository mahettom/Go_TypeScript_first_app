package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

// ————————————————————————————————————————— Give the structure of Todo endpoint
type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Body  int    `json:"body"`
}

func main() {
	fmt.Print("hello world")

	// ————————————————————————————————————————— Assigned to app whatever the type is
	app := fiber.New()

	// ————————————————————————————————————————— ?
	todos := []Todo{}

	// ————————————————————————————————————————— Make sure our app is up&running [postman]
	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		// ??
		todo := &Todo{}

		// ———————————————————————————————————— If there is a error, redirect or return the err
		if err := c.BodyParser(todo); err != nil {
			return err
		}

		todo.ID = len(todos) + 1

		// ——————————————————————————————————— Point at todo(of post) and append it in the todos list
		todos = append(todos, *todo)

		return c.JSON(todos)

	})

	log.Fatal(app.Listen(":4000"))
}
