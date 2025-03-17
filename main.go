package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct {
	ID        int  `json:"id"`
	Completed bool `json:"completed"`
	Body      string  `json:"body"`
}

func main() {
	fmt.Println("HELLO")
	app := fiber.New()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	todos := []Todo{}

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {

		todo := &Todo{} //save to the memory address

		// example of memory address and pointers
		// var x int = 5 //0x0001
		// var p *int = &x // points to the m. address of x ie. 0x001
		// fmt.Println(p) // 0x001
		// fmt.Println(*p) // prints the value at the memory address of p, ie. 5

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Body is required"})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo) //since todo is the memory address, we want to get the value out ot it

		return c.Status(201).JSON(todo)
	})

	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error{
		id := c.Params("id")
		updatedData := new(Todo)
		if err := c.BodyParser(updatedData);  err != nil {
			return err
		}
		for i, todo := range todos{
			if fmt.Sprint(todo.ID) == id {
				if updatedData.Body != ""{
					todos[i].Body = updatedData.Body
				}
				if updatedData.Completed {
					todos[i].Completed = updatedData.Completed
				}
				// todos[i].Completed = !todos[i].Completed
				todos[i].Completed = true
				// todos[i].Body =
				return c.Status(200).JSON(todos[i])
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	})

	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error{
		id := c.Params("id")
		for i, todo := range todos{
			if fmt.Sprint(todo.ID) == id {
				todos = append(todos[:i], todos[i+1:]...)
				//so here we are appending the todos from the start to the index before the one we want to delete and add the rest after the one we want to delete
				return c.Status(200).JSON(fiber.Map{"message": "Todo deleted"})
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	})

	log.Fatal(app.Listen(":" +PORT))
}
