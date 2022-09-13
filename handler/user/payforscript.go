package user

import (
	"uima/handler"
	"uima/model"

	"github.com/gin-gonic/gin"
)

type OrderId struct {
	Id int `json:"order_id"`
}

// @Summary "订单付款(剧本)"
// @Describtion "根据带订单进行更改订单的状态为“已付款”，默认为“未付款”"
// @Tags user
// @Accept json
// @Producer json
// @Param token header string true "token"
// @Param object body OrderId true "订单ID"
// @Success 200 "订单更新成功"
// @Failure 400 "输入格式错误"
// @Failure 401 "身份认证失效"
// @Failure 500 "服务器错误"
// @Router /user/payforscript [put]
func PayforScript(c *gin.Context) {
	var order_id OrderId
	if err := c.BindJSON(&order_id); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "输入格式有误",
		})
		return
	}

	if err := model.UpdateOrderStatus(order_id.Id); err != nil {
		handler.SendBadRequest(c, "订单更新失败，请重新付款！", nil, err)
		return
	}

	handler.SendResponse(c, "订单更新成功，您已付款！", nil)
}
