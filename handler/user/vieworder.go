package user

import (
	"uima/handler"
	"uima/model"

	"github.com/gin-gonic/gin"
)

// @Summary "查看订单"
// @Description "查看所有订单(含时间，价格，剧本名字)"
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Success 200 "成功"
// @Failure 401 "身份验证错误"
// @Failure 500 "服务器错误"
// @Router /user/vieworder [get]
func ViewOrder(c *gin.Context) {
	users_id := c.MustGet("id").(int)

	oderinfor, err := model.GetScriptOrder(users_id)
	if err != nil {
		handler.SendError(c, "返回订单数据失败！", nil, err)
		return
	}
	handler.SendResponse(c, "查询订单成功", oderinfor)
}
