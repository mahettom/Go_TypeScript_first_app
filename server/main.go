package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

// ————————————————————————————————————————————  Give the structure of Todo endpoint
type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Body  string `json:"body"`
}

func main() {
	fmt.Print("hello world")

	// ———————————————————————————————————————— Assign to app, whatever the type is
	app := fiber.New()

	// ————————————————————————————————————— ?? Assign a array of object with the struc Todo ??
	todos := []Todo{}

	// ————————————————————————————————————————— Make sure our app is up&running [postman]
	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// ————————————————————————————————————————————————————————————————————————————————————————
	// ————————————————————————————————————————————————————————————————————————————— POST ROUTE
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		// ??
		todo := &Todo{}

		// ———————————————————————————————————— If there is a error, redirect or return the err
		if err := c.BodyParser(todo); err != nil {
			return err
		}

		todo.ID = len(todos) + 1

		// ——————————————————————————————————— Point at todo(of post) & append it in the todos list
		todos = append(todos, *todo)

		return c.JSON(todos)
	})

	// ————————————————————————————————————————————————————————————————————————————————————————
	// ———————————————————————————————————————————————————————————————————————————— PATCH ROUTE
	app.Patch("api/todos/:id/done", func(c *fiber.Ctx) error {

		// ——————————————————————————————————— Access parameter & assign it to get ID or throw ID(err)
		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(401).SendString("Invalid ID")
		}

		// —————————————————————————————————— Find corresponding ID in todos & update Done to true
		for i, todo := range todos {
			if todo.ID == id {
				todos[i].Done = true
				break
			}
		}
		return c.JSON((todos))
	})

	// ————————————————————————————————————————————————————————————————————————————————————————
	// ———————————————————————————————————————————————————————————————————————————— PATCH ROUTE
	app.Get("/api/todos", func(c *fiber.Ctx) error {

		return c.JSON(todos)
	})

	log.Fatal(app.Listen(":4000"))
}
