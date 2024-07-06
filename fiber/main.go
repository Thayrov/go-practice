package main

import (
	"github.com/gofiber/fiber/v2"
)

// ToDo struct represents a task with an ID, title, and completion status
type ToDo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// todos is a slice that holds all ToDo items
var todos = []ToDo{}

func main() {
	// Create a new Fiber app instance
	app := fiber.New()

	// Define a GET route for the root URL
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Define a POST route to create a new ToDo item
	app.Post("/todos", func(c *fiber.Ctx) error {
		todo := new(ToDo)                          // Create a new ToDo instance
		if err := c.BodyParser(todo); err != nil { // Parse the request body into the ToDo instance
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
		}
		todo.ID = len(todos) + 1                        // Assign an ID to the new ToDo item
		todos = append(todos, *todo)                    // Append the new ToDo item to the todos slice
		return c.Status(fiber.StatusCreated).JSON(todo) // Return the created ToDo item as a JSON response
	})

	// Define a GET route to retrieve all ToDo items
	app.Get("/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos) // Return all ToDo items as a JSON response
	})

	// Define a GET route to retrieve a single ToDo item by its ID
	app.Get("/todos/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id") // Get the ID from the URL parameters and convert it to an integer
		if err != nil {              // Check for conversion errors
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
		}
		for _, todo := range todos { // Iterate through the todos slice
			if todo.ID == id { // Check if the ID matches
				return c.JSON(todo) // Return the matching ToDo item as a JSON response
			}
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "ToDo not found"}) // Return an error if the ID was not found
	})

	// Define a PUT route to update a ToDo item by its ID
	app.Put("/todos/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id") // Get the ID from the URL parameters and convert it to an integer
		if err != nil {              // Check for conversion errors
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
		}
		for i, todo := range todos { // Iterate through the todos slice
			if todo.ID == id { // Check if the ID matches
				if err := c.BodyParser(&todos[i]); err != nil { // Parse the request body into the matching ToDo item
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
				}
				todos[i].ID = id        // Ensure the ID remains unchanged
				return c.JSON(todos[i]) // Return the updated ToDo item as a JSON response
			}
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "ToDo not found"}) // Return an error if the ID was not found
	})

	// Define a DELETE route to remove a ToDo item by its ID
	app.Delete("/todos/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id") // Get the ID from the URL parameters and convert it to an integer
		if err != nil {              // Check for conversion errors
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID"})
		}
		for i, todo := range todos { // Iterate through the todos slice
			if todo.ID == id { // Check if the ID matches
				todos = append(todos[:i], todos[i+1:]...)  // Remove the matching ToDo item from the slice
				return c.SendStatus(fiber.StatusNoContent) // Return a No Content status
			}
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "ToDo not found"}) // Return an error if the ID was not found
	})

	// Start the Fiber app on port 3000
	app.Listen(":3000")
}
