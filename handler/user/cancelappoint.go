package user

import (
	"strconv"
	"uima/handler"
	"uima/model"

	"github.com/gin-gonic/gin"
)

type ScriptId struct {
	Id string `json:"script_id"` // 剧本的ID
}

// @Summary "取消预约剧本"
// @Description "将预约的剧本取消"
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Param object body ScriptId true "剧本Id"
// @Success 200 "取消预约成功"
// @Failure 400 "发送请求失败"
// @Failure 401 "身份验证失败"
// @Failure 500 "服务器错误，取消收藏失败"
// @Router /user/cancelappoint [post]
func CancelScriptAppoint(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	users_id := c.MustGet("id").(int)

	var scriptId ScriptId
	if err := c.BindJSON(&scriptId); err != nil {
		handler.SendBadRequest(c, "Lack Param or Param Not Satisfiable.", nil, err)
		return
	}

	scriptid, _ := strconv.Atoi(scriptId.Id)

	if err := model.CancelScript(users_id, scriptid); err != nil {
		handler.SendError(c, "取消预约失败!", nil, err)
		return
	}

	handler.SendResponse(c, "取消预约成功！", nil)
}
