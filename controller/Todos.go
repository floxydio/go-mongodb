package controller

import (
	"context"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gomongodb/database"
	"gomongodb/models"
)

func GetTodos(c echo.Context) error {
	collection := database.Client.Database("testmern").Collection("todos")
	dataTodos, errTodos := collection.Find(context.Background(), models.Todo{})
	if errTodos != nil {
		return c.JSON(500, echo.Map{
			"message": "Error when get data todos",
		})
	}
	var todos []models.Todo
	for dataTodos.Next(context.Background()) {
		var todo models.Todo
		err := dataTodos.Decode(&todo)
		if err != nil {
			return c.JSON(500, echo.Map{
				"message": "Error when decode data todos",
			})
		}
		todos = append(todos, todo)
	}

	return c.JSON(200, echo.Map{
		"message": "Success get data todos",
		"data":    todos,
	})
}

func CreateTodo(c echo.Context) error {
	collection := database.Client.Database("testmern").Collection("todos")
	var todo models.Todo
	if err := c.Bind(&todo); err != nil {
		return c.JSON(500, echo.Map{
			"message": "Error when bind data",
		})
	}
	_, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		return c.JSON(500, echo.Map{
			"message": "Error when insert data",
		})
	}
	return c.JSON(201, echo.Map{
		"message": "Success insert data",
	})
}

func UpdateTodo(c echo.Context) error {
	collection := database.Client.Database("testmern").Collection("todos")
	var todo models.Todo
	if err := c.Bind(&todo); err != nil {
		return c.JSON(500, echo.Map{
			"message": "Error when bind data",
		})
	}
	id := c.Param("id")
	hex, _ := primitive.ObjectIDFromHex(id)
	_, err := collection.UpdateOne(context.TODO(), bson.D{{"_id", hex}}, bson.D{{"$set", bson.D{{"title", todo.Title}, {"description", todo.Description}}}})
	if err != nil {
		return c.JSON(500, echo.Map{
			"message": "Error when update data",
		})
	}
	return c.JSON(201, echo.Map{
		"message": "Success update data",
	})
}

func DeleteTodos(c echo.Context) error {
	collection := database.Client.Database("testmern").Collection("todos")
	id := c.Param("id")
	hex, _ := primitive.ObjectIDFromHex(id)

	_, err := collection.DeleteOne(context.TODO(), bson.D{{"_id", hex}})

	if err != nil {
		return c.JSON(500, "Something Went Wrong")
	}

	return c.JSON(200, "Sukses Delet")
}