package main

import (
	"encoding/json" // Import for JSON encoding and decoding
	"fmt"           // Import for formatted I/O
	"net/http"      // Import for HTTP server
	"strconv"       // Import for string conversion utilities
	"strings"       // Import for string manipulation
)

// ToDo struct represents a task with an ID, title, and completion status
type ToDo struct {
	ID    int    `json:"id"`    // ID of the ToDo item
	Title string `json:"title"` // Title of the ToDo item
	Done  bool   `json:"done"`  // Completion status of the ToDo item
}

// todos is a slice that holds all ToDo items
var todos = []ToDo{}

func main() {
	http.HandleFunc("/", helloHandler)      // Handle requests to the root URL
	http.HandleFunc("/todos", todosHandler) // Handle requests to /todos for creating and listing ToDos
	http.HandleFunc("/todos/", todoHandler) // Handle requests to /todos/{id} for CRUD operations on a single ToDo

	fmt.Println("Server is listening on port 3000") // Print message to indicate the server is running
	http.ListenAndServe(":3000", nil)               // Start the HTTP server on port 3000
}

// helloHandler handles requests to the root URL
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!") // Write "Hello, World!" to the response
}

// todosHandler handles requests to /todos for creating and listing ToDos
func todosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTodos(w, r) // Handle GET request to retrieve all ToDos
	case http.MethodPost:
		createTodo(w, r) // Handle POST request to create a new ToDo
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed) // Return 405 error for unsupported methods
	}
}

// todoHandler handles requests to /todos/{id} for CRUD operations on a single ToDo
func todoHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/todos/") // Extract ID from the URL path
	id, err := strconv.Atoi(idStr)                     // Convert ID from string to integer
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest) // Return 400 error if ID is not valid
		return
	}

	switch r.Method {
	case http.MethodGet:
		getTodoByID(w, r, id) // Handle GET request to retrieve a single ToDo by ID
	case http.MethodPut:
		updateTodoByID(w, r, id) // Handle PUT request to update a ToDo by ID
	case http.MethodDelete:
		deleteTodoByID(w, r, id) // Handle DELETE request to remove a ToDo by ID
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed) // Return 405 error for unsupported methods
	}
}

// getTodos retrieves all ToDos and sends them as a JSON response
func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set response content type to JSON
	json.NewEncoder(w).Encode(todos)                   // Encode todos slice to JSON and write to response
}

// createTodo parses the request body to create a new ToDo item
func createTodo(w http.ResponseWriter, r *http.Request) {
	var todo ToDo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil { // Decode JSON request body into todo struct
		http.Error(w, "Cannot parse JSON", http.StatusBadRequest) // Return 400 error if JSON is invalid
		return
	}
	todo.ID = len(todos) + 1                           // Assign an ID to the new ToDo item
	todos = append(todos, todo)                        // Append the new ToDo item to the todos slice
	w.Header().Set("Content-Type", "application/json") // Set response content type to JSON
	w.WriteHeader(http.StatusCreated)                  // Set status code to 201 Created
	json.NewEncoder(w).Encode(todo)                    // Encode the created ToDo item to JSON and write to response
}

// getTodoByID retrieves a single ToDo item by its ID and sends it as a JSON response
func getTodoByID(w http.ResponseWriter, r *http.Request, id int) {
	for _, todo := range todos { // Iterate through the todos slice
		if todo.ID == id { // Check if the ID matches
			w.Header().Set("Content-Type", "application/json") // Set response content type to JSON
			json.NewEncoder(w).Encode(todo)                    // Encode the matching ToDo item to JSON and write to response
			return
		}
	}
	http.Error(w, "ToDo not found", http.StatusNotFound) // Return 404 error if the ToDo item is not found
}

// updateTodoByID parses the request body to update a ToDo item by its ID
func updateTodoByID(w http.ResponseWriter, r *http.Request, id int) {
	for i, todo := range todos { // Iterate through the todos slice
		if todo.ID == id { // Check if the ID matches
			if err := json.NewDecoder(r.Body).Decode(&todos[i]); err != nil { // Decode JSON request body into the matching ToDo item
				http.Error(w, "Cannot parse JSON", http.StatusBadRequest) // Return 400 error if JSON is invalid
				return
			}
			todos[i].ID = id                                   // Ensure the ID remains unchanged
			w.Header().Set("Content-Type", "application/json") // Set response content type to JSON
			json.NewEncoder(w).Encode(todos[i])                // Encode the updated ToDo item to JSON and write to response
			return
		}
	}
	http.Error(w, "ToDo not found", http.StatusNotFound) // Return 404 error if the ToDo item is not found
}

// deleteTodoByID removes a ToDo item by its ID
func deleteTodoByID(w http.ResponseWriter, r *http.Request, id int) {
	for i, todo := range todos { // Iterate through the todos slice
		if todo.ID == id { // Check if the ID matches
			todos = append(todos[:i], todos[i+1:]...) // Remove the matching ToDo item from the slice
			w.WriteHeader(http.StatusNoContent)       // Set status code to 204 No Content
			return
		}
	}
	http.Error(w, "ToDo not found", http.StatusNotFound) // Return 404 error if the ToDo item is not found
}
