package script

import (
	"uima/handler"
	"uima/model"

	"github.com/gin-gonic/gin"
)

// @Summary "剧本主界面"
// @Description "返回所有剧本的封面、名字、ID"
// @Tags script
// @Accept application/json
// @Produce application/json
// @Success 200 {object} []model.Script "获取成功"
// @Failure 404 "获取失败"
// @Router /script/interface [get]
func Interface(c *gin.Context) {
	Scriptsinfor, err := model.GetScriptCoverandNameandBreifIntro()
	if err != nil {
		handler.SendBadRequest(c,"搜索失败",nil,err)
		return
	}
	handler.SendResponse(c, "搜索成功", Scriptsinfor)
}
