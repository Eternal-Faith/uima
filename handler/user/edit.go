package user

import (
	"uima/handler"
	"uima/model"

	"github.com/gin-gonic/gin"
)

type UserInfor struct {
	Phone           string `gorm:"column:phone;type:varchar(11);unique;NOT NULL" json:"phone"` // 电话号码
	Nickname        string `gorm:"column:nickname;type:varchar(255)" json:"nickname"`          // 用户名
	Gender          string `gorm:"column:gender;type:varchar(255)" json:"gender"`              // 性别
	Email           string `gorm:"column:email;type:varchar(255)" json:"email"`                // 邮箱
	Realname        string `gorm:"column:realname;type:varchar(255)" json:"realname"`          // 真实姓名
	Idcard          string `gorm:"column:idcard;type:varchar(255)" json:"idcard"`              // 身份证号
}
// @Summary "设置个人信息"
// @Describtion "头像(单独写)，昵称，真实姓名，性别，身份证，手机，邮箱,密码(单独写)"
// @Tags user
// @Accept json
// @Producer json
// @Param token header string true "token"
// @Param object body UserInfor true "user"
// @Success 200 "登陆成功"
// @Failure 400 "输入格式错误"
// @Failure 404 "用户不存在"
// @Failure 401 "身份认证失效"
// @Failure 500 "服务器错误"
// @Router /user/edit [post]
func Edit(c *gin.Context) {
	id := c.MustGet("id")
	var userinfor model.User
	if err := c.BindJSON(&userinfor); err != nil {
		handler.SendBadRequest(c, "Lack Param or Param Not Satisfiable.",nil, err) 
		return
	}
	userinfor.Id = id.(int) //开始获得的id是接口类型

	if userinfor.Nickname == "" {
		handler.SendBadRequest(c, "昵称不可为空！",nil,nil) 
		return
	}

	for _, char := range userinfor.Nickname {//遍历字符串
		if string(char) == " " {
			handler.SendBadRequest(c, "昵称中不能含有空格！", nil,nil) 
			return
		}
	}

	if err := model.ChangeUserInfor(userinfor); err != nil {
		handler.SendBadRequest(c, "修改失败！",  nil,nil)
		return
	}

	handler.SendResponse(c, "修改成功!",nil)
}
