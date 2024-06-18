package ping

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PingRouter struct{}

func (PingRouter) SetupEndpoints(r *gin.Engine, db *gorm.DB) {
	r.GET("/ping", ping(db))
}

func ping(_ *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pon",
		})
	}
}
