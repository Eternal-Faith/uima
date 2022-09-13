package user

import (
	"uima/handler"
	"uima/model"

	"github.com/gin-gonic/gin"
)

type Status struct {
	Status  string `json:"status"` //标签
}

// @Summary "剧本分类"
// @Description "通过标签来进行分类"
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Param object body Status true "预约状态"
// @Success 200 "搜索成功"
// @Failure 400 "输入格式有误"
// @Failure 404 "搜索失败"
// @Failure 500 "服务器错误"
// @Router /user/searchappointment [post]
func SearchAppointment(c *gin.Context) {
	Id := c.MustGet("id")
	id := Id.(int)
	var a Status
	if err := c.BindJSON(&a); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "输入格式有误",
		})
		return
	}

	appointmentinfor, err := model.GetAppointmentByStatus(a.Status,id)

	if err != nil {
		handler.SendError(c, "搜素失败", nil, err)
		return
	}

	handler.SendResponse(c, "搜索成功", appointmentinfor)
}