package user

import (
	"uima/handler"
	"uima/model"

	"github.com/gin-gonic/gin"
)

// @Summary "预约剧本"
// @Description "预约时间，剧本id，剧本封面，预约状态"
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Param object body ScriptId true "剧本Id"
// @Success 200 "收藏成功"
// @Failure 401 "身份验证失败"
// @Failure 500 "服务器错误，收藏失败"
// @Router /user/appointment [post]
func Appointment(c *gin.Context) {

	users_id := c.MustGet("id").(int)

	var scriptId ScriptId
	if err := c.BindJSON(&scriptId); err != nil {
		handler.SendBadRequest(c, "Lack Param or Param Not Satisfiable.", nil, err)
		return
	}

	scriptInfor, err := model.GetScriptInfor(scriptId.Id)
	if err != nil {
		handler.SendError(c, "获取剧本数据失败!", nil, err)
		return
	}

	var appointment model.ScriptAppointments

	appointment.UsersId = users_id
	appointment.ScriptsCover = scriptInfor.Avatar
	appointment.ScriptsName = scriptInfor.ScriptName
	appointment.ScriptsId = scriptInfor.Id
	appointment.Status = "预约中"
	appointment.Time = "2022-05-26 15:04:00"

	if err := model.DB.Create(&appointment).Error; err != nil {
		handler.SendError(c, "预约失败", nil, err)
		return
	}

	handler.SendResponse(c, "预约剧本成功！", nil)
}
