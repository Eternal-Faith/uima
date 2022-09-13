package user

import (
	"strconv"
	"time"
	"uima/handler"
	"uima/model"

	"github.com/gin-gonic/gin"
)

// @Summary "创建订单(剧本)"
// @Describtion "订单包含用户电话,订单号,价格,订单图片,创建订单时间,付款时间,订单内容"
// @Tags user
// @Accept json
// @Producer json
// @Param token header string true "token"
// @Param object body ScriptId true "剧本Id"
// @Success 200 "创建订单成功"
// @Failure 400 "输入格式有误"
// @Failure 500 "服务器错误"
// @Router /user/order [post]
func CreateScriptOrder(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	
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

	var scriptorder model.ScriptOrders
	scriptorder.UserId = users_id
	scriptorder.ScriptId, _ = strconv.Atoi(scriptId.Id)
	scriptorder.Avatar = scriptInfor.Avatar
	scriptorder.Createtime = time.Now().Format("2006-01-02 15:04:00")
	scriptorder.Type = "剧本订单"
	scriptorder.Information = scriptInfor.BriefIntro
	scriptorder.Paymenttime = "你还没有支付！"
	scriptorder.Price = scriptInfor.Price
	scriptorder.Status = "待付款"
	scriptorder.ScriptName = scriptInfor.ScriptName

	if err := model.DB.Create(&scriptorder).Error; err != nil {
		handler.SendError(c, "创建订单失败!", nil, err)
		return
	}

	handler.SendResponse(c, "创建订单成功!", nil)
}
