package shop

import (
	"uima/model"

	"github.com/gin-gonic/gin"
)

// @Summary "单个商店信息"
// @Tags shop
// @Description "获取单个商店详细信息"
// @Accept json
// @Produce json
// @Param shop_id path string true "shop_id"
// @Success 200 "获取成功"
// @Failure 401 "获取参数失败"
// @Failure 400 "获取失败"
// @Router /shop/single_shop/:shop_id [get]
func GetSingleShopInfo(c *gin.Context) {
	id := c.Param("shop_id")

	shop, err := model.GetSingleShopInfo(id)
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
