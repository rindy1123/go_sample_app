package main

import (
	"os"

	"github.com/api/internal/infra"
	"github.com/api/internal/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const PORT = ":8080"

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{os.Getenv("FRONT_ORIGIN")},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))
	db := infra.SetupDB()
	routes.SetupEndpoints(r, db)
	r.Run(PORT)
}
