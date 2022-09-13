package broadcast

import (
	"fmt"
	"uima/model"

	"github.com/gin-gonic/gin"
)

// @Summary "获取单个精彩放送信息"
// @Tags bro
// @Description "获取精彩放送详细信息"
// @Accept json
// @Produce json
// @Param broadcast_id path string true "broadcast_id"
// @Success 200 "获取成功"
// @Failure 401 "身份验证失败"
// @Failure 400 "获取失败"
// @Router /broadcast/get_single/:broadcast_id [get]
func GetSingleBroadcast(c *gin.Context) {
	id := c.Param("broadcast_id")
	// var bro model.Broadcast
	fmt.Println(id)
	bro, err := model.GetSingleBroadcastInfo(id)

	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "获取失败",
		})
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    bro,
	})
}
