package auth

import (
	"fmt"
	"strconv"
	"uima/model"
	"uima/pkg/tokens"

	"github.com/gin-gonic/gin"
)

type UserRegister struct {
	Phone string `json:"phone"` 
	Password string `json:"password"`
	Confirmpassword string `json:"confirm_password"`
}
// @Summary "注册"
// @Description "注册一个新用户"
// @tags auth
// @Accept json
// @Produce json
// @Param user body UserRegister true "user"
// @Success 200 "用户创建成功"
// @Failure 400 "输入有误，格式错误"
// @Failure 401 "电话号码重复"
// @Router /auth/register [post]
func Register(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "输入有误，格式错误"})
		return
	}
	//电话位数问题前端处理
	fmt.Println(user.Phone)
	if user.Password != user.ConfirmPassword {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "两次密码输入不一致，请重新输入"})
		return
	}
	if _, a := model.IfExistUserPhone(user.Phone); a != 1 {
		c.JSON(401, gin.H{
			"code":    401,
			"message": "对不起，该电话号码已经被绑定",
		})
		return
	}
	user.Password = model.GeneratePasswordHash(user.Password)
	user_id := model.Register(user.Phone, user.Password)

	fmt.Println(user.Phone, user_id)
	user.Id, _ = strconv.Atoi(user_id)
	signedToken := tokens.GenerateToken(user.Id)
	c.JSON(200, gin.H{
		"code":    200,
		"message": "将用户id作为token保留",
		"data":    signedToken,
	})
}
