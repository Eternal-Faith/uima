package broadcast

import (
	"uima/model"

	"github.com/gin-gonic/gin"
)

// @Summary "精彩放送"
// @Tags bro
// @Description "精彩放送"
// @Accept json
// @Produce json
// @Param broadcast body model.Broadcast true "Broadcast"
// @Success 200 "上传成功"
// @Failure 401 "身份验证失败"
// @Failure 400 "上传失败"
// @Router /broadcast/get_all [get]
func GetBroadcast(c *gin.Context) {

	// id := c.MustGet("id").(int)
	// fmt.Println(id)
	Broadcast, err := model.GetBroadcastInfo()
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "获取失败",
		})
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    Broadcast,
	})
}
