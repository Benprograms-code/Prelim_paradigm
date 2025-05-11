package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// json struct to store data
type todo struct {
	ID        string `json:"id"`
	Item      string `json:"title"`
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

// json converter function
func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.Run("Localhost:9090")
}
