package routers

import (
	"github.com/CareyWang/gofun/handles"
	"github.com/gin-gonic/gin"
)

func InitApiRouter(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.GET("/ping", func(context *gin.Context) {
			context.String(200, "pong!")
		})

		v1.GET("/id", handles.Generate)
	}
}
