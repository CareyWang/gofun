package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"

	_ "gofun/docs"
)

func InitDocsRouter(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
