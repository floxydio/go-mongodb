package main

import (
	"gomongodb/controller"
	"gomongodb/database"
	"github.com/labstack/echo/v4"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()

	database.ConnectDatabase()

	// Todos
	e.GET("/todos", controller.GetTodos)
	e.POST("/todos", controller.CreateTodo)
	e.PUT("/update-todos/:id", controller.UpdateTodo)
	e.DELETE("/delete-todos/:id", controller.DeleteTodos)

	// Users

	e.Logger.Fatal(e.Start(":3000"))

}