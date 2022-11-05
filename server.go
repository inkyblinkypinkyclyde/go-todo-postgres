package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func indexHandler(c *fiber.Ctx, db *sql.DB) error {
	var res string
	var todos []string
	rows, err := db.Query("SELECT * FROM todos") // Query the database
	defer rows.Close()                           // Close the statement when we leave main() / the program terminates
	if err != nil {                              // If there is an error
		log.Fatalln(err)           // Log the error
		c.JSON("An error occured") // Return a 500 error
	}
	for rows.Next() { // Iterate through the rows
		rows.Scan(&res)            // Get the value from the column with the name "item"
		todos = append(todos, res) // Append the value to the slice
	}
	return c.Render("index", fiber.Map{ // Render the index template
		"Todos": todos, // Pass the todos slice to the template
	})
}

type todo struct { // Create a struct to hold the todo item
	Item string `json:"item"` // The item field
}

func postHandler(c *fiber.Ctx, db *sql.DB) error { // Create a handler function which takes a context and a database connection
	newTodo := todo{}                              // Create a new todo struct
	if err := c.BodyParser(&newTodo); err != nil { // Parse the request body into the newTodo struct
		log.Printf("An error occured: %v", err) // Log the error
		return c.SendString(err.Error())        // Return the error
	}
	fmt.Printf("%v", newTodo) // Print the newTodo struct
	if newTodo.Item != "" {   // If the item field is not empty
		_, err := db.Exec("INSERT into todos VALUES ($1)", newTodo.Item) // Insert the item into the database
		if err != nil {                                                  // If there is an error
			log.Fatalf("An error occured while executing query: %v", err) // Log the error
		}
	}

	return c.Redirect("/") // Redirect to the index page
}

func putHandler(c *fiber.Ctx, db *sql.DB) error { // Create a handler function which takes a context and a database connection
	olditem := c.Query("olditem")                                       // Get the old item from the query string
	newitem := c.Query("newitem")                                       // Get the new item from the query string
	db.Exec("UPDATE todos SET item=$1 WHERE item=$2", newitem, olditem) // Update the item in the database
	return c.Redirect("/")                                              // Redirect to the index page
}

func deleteHandler(c *fiber.Ctx, db *sql.DB) error { // Create a handler function which takes a context and a database connection
	todoToDelete := c.Query("item")                          // Get the item from the query string
	db.Exec("DELETE from todos WHERE item=$1", todoToDelete) // Delete the item from the database
	return c.SendString("deleted")                           // Return a 200 OK response
}

func main() { // Our main function
	connStr := "postgresql://localhost/go-todo?user=richardgannon&password=postgres&sslmode=disable" // Our connection string
	db, err := sql.Open("postgres", connStr)                                                         // Open a database connection
	if err != nil {                                                                                  // If there is an error
		log.Fatal(err) // Log the error
	}

	engine := html.New("./views", ".html") // Create a new HTML engine
	app := fiber.New(fiber.Config{         // Create a new Fiber instance
		Views: engine, // Set the views engine
	})

	app.Get("/", func(c *fiber.Ctx) error { // Create a route to handle GET requests to the index page
		return indexHandler(c, db)
	})

	app.Post("/", func(c *fiber.Ctx) error { // Create a route to handle POST requests to the index page
		return postHandler(c, db)
	})

	app.Put("/update", func(c *fiber.Ctx) error { // Create a route to handle PUT requests to the update page
		return putHandler(c, db)
	})

	app.Delete("/delete", func(c *fiber.Ctx) error { // Create a route to handle DELETE requests to the delete page
		return deleteHandler(c, db)
	})

	port := os.Getenv("PORT") // Get the port from the environment
	if port == "" {           // If the port is empty
		port = "3000" // Set the port to 3000
	}

	app.Static("/", "./public")                       // Serve static files from the public directory
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port))) // Start the server

}
