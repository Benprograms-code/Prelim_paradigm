package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// json struct to store data
type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

// list of example tasks
var todos = []todo{
	{ID: "1", Item: "Learn Go", Completed: false},
	{ID: "2", Item: "Build a web app", Completed: false},
	{ID: "3", Item: "Deploy to production", Completed: false},
	{ID: "4", Item: "Learn Docker", Completed: false},
	{ID: "5", Item: "Learn Kubernetes", Completed: false},
	{ID: "6", Item: "Learn AWS", Completed: false},
	{ID: "7", Item: "Learn Azure", Completed: false},
	{ID: "8", Item: "Learn GCP", Completed: false},
	{ID: "9", Item: "Learn Terraform", Completed: false},
	{ID: "10", Item: "Learn Ansible", Completed: false},
	{ID: "11", Item: "Learn CI/CD", Completed: false},
}

func addTodo(context *gin.Context) {
	var newtodo todo

	//error catchment using golang error handling
	if error := context.BindJSON(&newtodo); error != nil {
		return
	}

	todos = append(todos, newtodo)
	context.IndentedJSON(http.StatusCreated, newtodo)

}

func getTodo(context *gin.Context) {
	id := context.Param("id")

	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

// json converter function
func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}
func toggleTodoSatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}

	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoSatus)
	router.POST("/todos", addTodo)
	router.Run("localhost:9090")
}
