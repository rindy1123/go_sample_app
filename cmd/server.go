package main

import (
	"github.com/api/internal/models"
	"github.com/api/internal/routes"
	"github.com/gin-gonic/gin"
)

const PORT = ":8080"

func main() {
	r := gin.Default()
	db := models.SetupDB()
	routes.SetupEndpoints(r, db)
	r.Run(PORT)
}
