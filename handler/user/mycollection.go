package user

import (
	"uima/handler"
	"uima/model"

	"github.com/gin-gonic/gin"
)

// @Summary "我的收藏(剧本)"
// @Description "获取用户的所有剧本收藏"
// @Tags user
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 "查看成功"
// @Failure 400 "发送请求失败"
// @Failure 401 "身份认证失效"
// @Failure 500 "服务器错误"
// @Router /user/mycollection [get]
func MyCollection(c *gin.Context) {
	Id := c.MustGet("id")
	id := Id.(int)

	Collections, err := model.GetCollectScript(id)
	if err != nil {
		handler.SendError(c, "查看失败！", nil, err)
		return
	}

	handler.SendResponse(c, "查看成功！", Collections)
}
