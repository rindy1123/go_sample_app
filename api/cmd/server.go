package main

import (
	"github.com/api/internal/routes"
	"github.com/api/internal/utils"
	"github.com/gin-gonic/gin"
)

const PORT = ":8080"

func main() {
	r := gin.Default()
	db := utils.SetupDB()
	routes.SetupEndpoints(r, db)
	r.Run(PORT)
}
