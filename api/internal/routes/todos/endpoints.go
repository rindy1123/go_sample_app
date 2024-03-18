package todos

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TodoRouter struct{}

func (TodoRouter) SetupEndpoints(r *gin.Engine, _ *gorm.DB) {
	r.GET("/todos", getTodos)
	r.GET("/todos/:id", getTodo)
	// r.POST("/todos", createTodo)
	// r.PUT("/todos/:id", updateTodo)
	// r.DELETE("/todos/:id", deleteTodo)
}

type todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos = []todo{
	{ID: 1, Title: "pay phone bill", Status: "active"},
	{ID: 2, Title: "buy some milk", Status: "active"},
	{ID: 3, Title: "learn golang", Status: "active"},
}

func getTodos(c *gin.Context) {
	c.JSON(200, todos)
}

func getTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid id"})
		return
	}
	for _, todo := range todos {
		if todo.ID == id {
			c.JSON(200, todo)
			return
		}
	}
}
