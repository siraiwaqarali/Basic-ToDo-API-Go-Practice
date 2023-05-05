package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ToDo struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos []ToDo = []ToDo{
	{ID: 1, Title: "Buy Milk", Completed: false},
	{ID: 2, Title: "Buy Eggs", Completed: false},
	{ID: 3, Title: "Buy Bread", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo ToDo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoById(id uint) (*ToDo, error) {
	for index, todo := range todos {
		if todo.ID == id {
			return &todos[index], nil
		}
	}

	return nil, errors.New("todo not found")
}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")

	idUint64, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "id is not valid",
		})
		return
	}

	todo, err := getTodoById(uint(idUint64))

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "todo not found",
		})
		return
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}

func getTodo(context *gin.Context) {
	id := context.Param("id")

	idUint64, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "id is not valid",
		})
		return
	}

	todo, err := getTodoById(uint(idUint64))
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "todo not found",
		})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func main() {
	router := gin.Default()

	router.GET("/api/todos", getTodos)
	router.POST("/api/add-todo", addTodo)
	router.PATCH("/api/toggle-todo-status/:id", toggleTodoStatus)
	router.GET("/api/todos/:id", getTodo)

	router.Run("localhost:3000")
}
