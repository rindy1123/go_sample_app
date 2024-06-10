package routes

import (
	"gorm.io/gorm"

	"github.com/api/internal/routes/ping"
	"github.com/api/internal/routes/todos"
	"github.com/gin-gonic/gin"
)

type Router interface {
	SetupEndpoints(r *gin.Engine, db *gorm.DB)
}

var routes = []Router{
	todos.TodoRouter{},
	ping.PingRouter{},
}

func SetupEndpoints(r *gin.Engine, db *gorm.DB) {
	for _, route := range routes {
		route.SetupEndpoints(r, db)
	}
}
