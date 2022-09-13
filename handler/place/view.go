package place

import (
	"uima/handler"
	"uima/model"

	"github.com/gin-gonic/gin"
)

type PlaceName struct {
	PlaceName string `gorm:"column:name;type:varchar(255);NOT NULL" json:"place_name"` // 地点名称
}

// @Summary "查看地点"
// @Description "查看地点信息"
// @Tags place
// @Accept application/json
// @Produce application/json
// @Param object body PlaceName true "地点名称"
// @Success 200 {object} model.Place
// @Failure 500 "服务器错误"
// @Router /place/view [post]
func ViewPlace(c *gin.Context) {
	var name PlaceName
	if err := c.BindJSON(&name); err != nil {
		handler.SendError(c, "Lack Param or Param Not Satisfiable.", nil, err)
		return
	}

	placeInfor, err := model.GetPlaceInfor(name.PlaceName)

	if err != nil {
		handler.SendBadRequest(c, "获取地点信息失败", nil, err)
		return
	}

	handler.SendResponse(c, "获取地点信息成功", placeInfor)

}
