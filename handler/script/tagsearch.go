package script

import (
	"uima/handler"
	"uima/model"

	"github.com/gin-gonic/gin"
)

type Tag struct {
	Tags string `json:"tag"` //标签
}

// @Summary "剧本分类"
// @Description "通过标签来进行分类"
// @Tags script
// @Accept application/json
// @Produce application/json
// @Param object body Tag true "标签"
// @Success 200 "搜索成功"
// @Failure 400 "输入格式有误"
// @Failure 404 "搜索失败"
// @Failure 500 "服务器错误"
// @Router /script/tag [post]
func TagSearch(c *gin.Context) {
	var a Tag
	if err := c.BindJSON(&a); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "输入格式有误",
		})
		return
	}

	scriptinfor, err := model.GetInforbyTag(a.Tags)

	if err != nil {
		handler.SendError(c, "搜素失败", nil, err)
		return
	}

	handler.SendResponse(c, "搜索成功", scriptinfor)
}
