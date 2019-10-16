package routers

import "github.com/gin-gonic/gin"

func Init(r *gin.Engine) {
	InitApiRouter(r)
	InitDocsRouter(r)
}
