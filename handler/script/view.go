package script

import (
	"uima/handler"
	"uima/model"

	"github.com/gin-gonic/gin"
)

type ScriptId struct {
	Id string `json:"script_id"` // 剧本的ID
}

// @Summary "查看剧本"
// @Description "查看剧本的具体内容"
// @Tags script
// @Accept application/json
// @Produce application/json
// @Param object body ScriptId true "剧本Id"
// @Success 200 "成功"
// @Failure 401 "身份验证错误"
// @Failure 500 "服务器错误"
// @Router /script/view [post]
func ViewScript(c *gin.Context) {
	var scriptId ScriptId
	if err := c.BindJSON(&scriptId); err != nil {
		handler.SendBadRequest(c, "Lack Param or Param Not Satisfiable.", nil, err)
		return
	}

	scriptInfo, err := model.GetScriptInfor(scriptId.Id)
	if err != nil {
		handler.SendError(c, "查看剧本失败", nil, err)
		return
	}

	handler.SendResponse(c, "查看剧本成功", scriptInfo)
}
