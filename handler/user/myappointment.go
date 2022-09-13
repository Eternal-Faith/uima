package user

import (
	"uima/handler"
	"uima/model"

	"github.com/gin-gonic/gin"
)

// @Summary "我的预约(剧本)"
// @Description "获取用户的所有剧本预约"
// @Tags user
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 "查看成功"
// @Failure 400 "发送请求失败"
// @Failure 401 "身份认证失效"
// @Failure 500 "服务器错误"
// @Router /user/myappointment [get]
func MyAppointment(c *gin.Context) {
	Id := c.MustGet("id")
	id := Id.(int)

	Appointments, err := model.GetScriptAppointment(id)
	if err != nil {
		handler.SendError(c, "查看预约失败！", nil, err)
		return
	}

	handler.SendResponse(c, "查看预约成功！", Appointments)
}