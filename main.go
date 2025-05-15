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
var to_dos = []todo{
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

// WEAKNESS 2 lack of OOP inheritance
// WorkTodo embeds todo and adds Priority
type WorkTodo struct {
	todo
	Priority string `json:"priority"`
}

// HomeTodo embeds todo and adds Location
type HomeTodo struct {
	todo
	Location string `json:"location"`
}

// Sample to_dos for both categories
var workto_dos = []WorkTodo{
	{todo: todo{ID: "1", Item: "Finish report", Completed: false}, Priority: "High"},
	{todo: todo{ID: "2", Item: "Email client", Completed: false}, Priority: "Medium"},
}

var hometo_dos = []HomeTodo{
	{todo: todo{ID: "3", Item: "Buy groceries", Completed: false}, Location: "Supermarket"},
	{todo: todo{ID: "4", Item: "Clean kitchen", Completed: false}, Location: "Home"},
}

func getWorkto_dos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, workto_dos)
}

func getHometo_dos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, hometo_dos)
}

// ____________________________________________________________________________________________________________________________________________________
func TodoAdd(context *gin.Context) {
	var newtodo todo

	if err := context.BindJSON(&newtodo); err != nil {
		//WEAKNESS 1: LACK OF ERROR HANDLING
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		//_____________________________________________________________________________________________________________________
		return
	}

	to_dos = append(to_dos, newtodo)
	context.IndentedJSON(http.StatusCreated, newtodo)

	//STRENGTH 1:GOROUTINES
	//async task to simulate logging
	go func(t todo) {
		println("Logged new todo asynchronously:", t.Item)
	}(newtodo)
	//_____________________________________________________________________________________________________________________
}

func getTodo(context *gin.Context) {
	id := context.Param("id")

	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
} //json converter function

func getto_dos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, to_dos)
}

func getTodoById(id string) (*todo, error) {
	for i, t := range to_dos {
		if t.ID == id {
			return &to_dos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}
func toggleto_dosatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}

	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}

// STRENGTH 2: GIN FRAMEWORK + native HTTP
func main() {
	router := gin.Default()
	router.GET("/to_dos", getto_dos)
	router.GET("/to_dos/:id", getTodo)
	router.PATCH("/to_dos/:id", toggleto_dosatus)
	router.POST("/to_dos", TodoAdd)

	router.GET("/to_dos/workto_dos", getWorkto_dos)
	router.GET("/to_dos/hometo_dos", getHometo_dos)

	router.Run("localhost:1234")
}

// ____________________________________________________________________________________________________________________________________________________
