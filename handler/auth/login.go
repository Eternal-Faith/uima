package auth

import (
	"fmt"
	"log"
	"net/http"
	"uima/model"
	"uima/pkg/tokens"
	"github.com/gin-gonic/gin"
)

type UserLogin struct {
	Phone string `json:"phone"` 
	Password string `json:"password"`
}

// @Summary "登录"
// @Describtion "输入电话密码验证用户信息实现登入"
// @Tags auth
// @Accept json
// @Producer json
// @Param user body UserLogin true "user"
// @Success 200 "登陆成功"
// @Failure 400 "输入格式错误"
// @Failure 404 "用户不存在"
// @Failure 401 "密码错误"
// @Router /auth/login [post]
func Login(c *gin.Context) {
	var user1 model.User
	
	if err := c.BindJSON(&user1); err != nil {
		c.JSON(400, gin.H{
			"code":400,
			"message":"输入格式有误",
		})
		log.Println(err)
		return
	}

	fmt.Println(user1.Phone, user1.Password)
	//验证用户是否存在（电话是否已经注册）
	if model.VerifyPhone(user1.Phone) != false {
		c.JSON(404, gin.H{
			"code":404,
			"message":"用户不存在",
		})
		return
	}

	user2, err := model.GetUserInfoByPhone(user1.Phone)
	if err != nil {
		log.Println(err)
	}

	//验证密码,密码可能重复所以还要电话（用这两个验证是否有这条数据在）
	if model.CheckPassword(user1.Password, user2.Password) == false {
		c.JSON(http.StatusUnauthorized, "password or account is wrong.")
		return
	} else {
		c.JSON(200, gin.H{
			"code":200,
			"message":"登陆成功,请将token放到请求头中",
			"token":tokens.GenerateToken(user2.Id),
		})
		return
	}

}
