package user

import (
	"strconv"
	"uima/handler"
	"uima/model"

	"github.com/gin-gonic/gin"
)

// @Summary "取消收藏剧本"
// @Description "将收藏的剧本取消"
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Param object body ScriptId true "剧本Id"
// @Success 200 "取消收藏成功"
// @Failure 400 "发送请求失败"
// @Failure 401 "身份验证失败"
// @Failure 500 "服务器错误，取消收藏收藏失败"
// @Router /user/cancel [post]
func CancelScriptCollection(c *gin.Context) {

	users_id := c.MustGet("id").(int)

	var scriptId ScriptId
	if err := c.BindJSON(&scriptId); err != nil {
		handler.SendBadRequest(c, "Lack Param or Param Not Satisfiable.", nil, err)
		return
	}

	scriptid, _ := strconv.Atoi(scriptId.Id)

	if err := model.CancelScript(users_id, scriptid); err != nil {
		handler.SendError(c, "取消收藏失败!", nil, err)
		return
	}

	handler.SendResponse(c, "取消收藏成功！", nil)
}
