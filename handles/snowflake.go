package handles

import (
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
)

// @Summary 获取全局唯一ID
// @Id 1
// @Tags snowflake
// @version 1.0
// @success 200 string string "成功返回全局唯一ID"
// @Router /v1/id [get]
func GenerateUniqID(context *gin.Context) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		context.JSON(500, err)
		return
	}

	context.JSON(200, node.Generate().Int64())
}
