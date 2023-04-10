package router

import (
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.POST("/oracle", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "oracle",
		})
	})

	return r
}
