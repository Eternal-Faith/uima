package user

import (
	"uima/handler"
	"uima/model"

	"github.com/gin-gonic/gin"
)

// @Summary "游览个人信息"
// @Description "获取用户的基本信息"
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Success 200 "游览信息成功"
// @Failure 401 "身份验证失败"
// @Failure 404 "客户端错误"
// @Router /user/view [post]
func Userinfo(c *gin.Context) {

	id := c.MustGet("id").(int)
	Userinformation, err := model.GetUserInfo(id)

	if err != nil {
		handler.SendBadRequest(c,"服务器错误",nil,err)
		return
	}

	handler.SendResponse(c, "搜索成功", Userinformation) 

}