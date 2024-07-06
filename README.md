# Go Practice

This project contains two versions of a ToDo list backend API implemented in Go: one using the Fiber framework and the other using the standard library.

## Project Structure

```
go-practice/
├── fiber/
│ ├── go.mod
│ ├── go.sum
│ └── main.go
├── stdlib/
│ ├── go.mod
│ └── main.go
└── README.md
```

## Running the Fiber Version

1. Navigate to the `fiber` directory:

   ```sh
   cd fiber
   ```

2. Install dependencies and run the server:

   ```sh
   go mod tidy
   go run main.go
   ```

3. The server will be running on http://localhost:3000.

## Running the Standard Library Version

1. Navigate to the stdlib directory:

   ```sh
   cd stdlib
   ```

2. Install dependencies and run the server:

   ```sh
   go mod tidy
   go run main.go
   ```

3. The server will be running on http://localhost:3000.
