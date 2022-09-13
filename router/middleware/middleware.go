package middleware

import (
	"uima/pkg/tokens"
	"net/http"

	"github.com/gin-gonic/gin"
)

//直接用在路径里，在每个api开始前使用一次，如果认证失败就直接报错
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		id, err := tokens.VerifyToken(token)

		c.Set("id", id)
		if err != nil || id == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Token invalid."})
			c.Abort()
			return
		}
	}
}
