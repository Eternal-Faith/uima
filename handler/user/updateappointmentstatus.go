package user

import (
	"uima/handler"
	"uima/model"

	"github.com/gin-gonic/gin"
)

type Appointments struct {
	Id     int    `json:"appointment_id"`
	Status string `json:"status"`
}

// @Summary "更新预约状态"
// @Describtion "预约的ID和修改的状态"
// @Tags user
// @Accept json
// @Producer json
// @Param token header string true "token"
// @Param object body Appointments true "预约ID及状态"
// @Success 200 "预约更新成功"
// @Failure 400 "输入格式错误"
// @Failure 401 "身份认证失效"
// @Failure 500 "服务器错误"
// @Router /user/updateappointmentstatus [put]
func UpdateAppointmentStatus(c *gin.Context) {
	var appointment Appointments
	if err := c.BindJSON(&appointment); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "输入格式有误",
		})
		return
	}

	if err := model.UpdateAppointmentStatus(appointment.Id, appointment.Status); err != nil {
		handler.SendBadRequest(c, "预约更新状态失败！", nil, err)
		return
	}

	handler.SendResponse(c, "预约更新状态成功！", nil)
}
