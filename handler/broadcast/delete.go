package broadcast

import (
	"uima/model"

	"github.com/gin-gonic/gin"
)

// @Summary "删除一个精彩放送信息"
// @Tags bro
// @Description "删除一个精彩放送的信息"
// @Accept json
// @Produce json
// @Param broadcast_id path string true "broadcast_id"
// @Success 200 "删除成功"
// @Failure 400 "删除失败"
// @Router /broadcast/:broadcast_id [delete]
func DeleteBroadcast(c *gin.Context) {
	id := c.Param("broadcast_id")
	if err := model.DeleteSingleBroInfo(id); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "删除失败",
		})
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "删除成功",
	})
}
