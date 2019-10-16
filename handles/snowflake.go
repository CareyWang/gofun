package handles

import (
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
)

func GenerateUniqID(context *gin.Context) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		context.JSON(500, err)
		return
	}

	context.JSON(200, node.Generate().Int64())
}
