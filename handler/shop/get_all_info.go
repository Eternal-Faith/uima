package shop

import (
	"uima/model"

	"github.com/gin-gonic/gin"
)

// @Summary "商店信息"
// @Tags shop
// @Description "获取商店详细信息"
// @Accept json
// @Produce json
// @Param shop body model.Shop true "shop"
// @Success 200 "上传成功"
// @Failure 401 "身份验证失败"
// @Failure 400 "上传失败"
// @Router /shop/get_all [get]
func GetAllShopInfo(c *gin.Context) {

	shop, err := model.GetShopInfo()
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "获取失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "获取成功",
		"data":    shop,
	})
}
