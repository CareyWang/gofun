package routes

import (
	"github.com/gin-gonic/gin"
	"gofun/handles"
)

func InitApiRouter(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.GET("/ping", func(context *gin.Context) {
			context.String(200, "pong!")
		})

		v1.GET("/id", handles.GenerateUniqID)

		v1.GET("/yoyu/:url", handles.GenerateSurge)
	}
}
