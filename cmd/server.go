package main

import (
	"github.com/api/internal/infra"
	"github.com/api/internal/routes"

	"github.com/gin-gonic/gin"
)

const PORT = ":8080"

func main() {
	r := gin.Default()
	db := infra.SetupDB()
	routes.SetupEndpoints(r, db)
	r.Run(PORT)
}
