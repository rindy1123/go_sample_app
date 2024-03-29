package todos

import (
	"net/http"

	"github.com/api/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type TodoRouter struct{}

func (TodoRouter) SetupEndpoints(r *gin.Engine, db *gorm.DB) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("todostatus", todoStatusValidator)
	}

	r.GET("/todos", list(db))
	r.GET("/todos/:id", get(db))
	r.POST("/todos", create(db))
	r.PUT("/todos/:id", update(db))
	r.DELETE("/todos/:id", delete(db))
}

func list(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		todos := models.Todo{}.List(db)
		c.JSON(http.StatusOK, toMultiple(todos))
	}
}

func get(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		todo := models.Todo{}.Get(db, id)
		c.JSON(http.StatusOK, toSingular(todo))
	}
}

func create(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var body PostReqBody
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id := models.Todo{Title: body.Title, Status: body.Status}.Create(db)
		c.JSON(http.StatusOK, gin.H{"id": id})
	}
}

func update(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var body UpdateReqBody
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id := c.Param("id")
		models.Todo{Title: body.Title, Status: body.Status}.Update(db, id)
		c.Status(http.StatusOK)
	}
}

func delete(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		models.Todo{}.HardDelete(db, id)
		c.Status(http.StatusOK)
	}
}
