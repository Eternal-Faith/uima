package broadcast

import (
	"time"
	"uima/handler"
	"uima/model"

	"github.com/gin-gonic/gin"
)

// @Summary "添加精彩放送信息"
// @Tags bro
// @Description "添加精彩放送"
// @Accept json
// @Produce json
// @Param broadcast body model.Broadcast true "broadcast"
// @Success 200 "上传成功"
// @Failure 401 "身份验证失败"
// @Failure 400 "上传失败"
// @Router /broadcast/basic_info [post]
func CreateBroadcast(c *gin.Context) {

	var bro model.Broadcast

	if err := c.BindJSON(&bro); err != nil {
		handler.SendBadRequest(c, "Lack Param or Param Not Satisfiable.", nil, err)
		return
	}
	bro.CreateTime = time.Now().Format("2006/01/02 15:04")

	if err := model.DB.Table("broadcasts").Create(&bro).Error; err != nil {
		handler.SendError(c, "添加精彩放送信息失败", nil, err)
		return
	}

	c.JSON(200, gin.H{
		"code":         200,
		"message":      "创建成功",
		"broadcast_id": bro.Id,
	})
}
