package place

import (
	"uima/handler"
	"uima/model"

	"github.com/gin-gonic/gin"
)

type Places struct {
	Name          string `gorm:"column:name;type:varchar(255);NOT NULL" json:"name"`                   // 地点名称
	Data          string `gorm:"column:data;type:varchar(1000);NOT NULL" json:"data"`                  // 地点资料
	Area          string `gorm:"column:area;type:varchar(255);NOT NULL" json:"area"`                   // 地点面积
	Visitor       string `gorm:"column:visitor;type:varchar(255);NOT NULL" json:"visitor"`             // 最大游客量
	Entertainment string `gorm:"column:entertainment;type:varchar(255);NOT NULL" json:"entertainment"` // 娱乐项目数量
	ScenicSpot    string `gorm:"column:scenic_spot;type:varchar(255);NOT NULL" json:"scenic_spot"`     // 特色景点数量
}

// @Summary "创建地点信息"
// @Describtion "上传地点的名称，介绍"
// @Tags place
// @Accept json
// @Producer json
// @Param object body Places true "地点"
// @Success 200 "创建成功"
// @Failure 500 "服务器错误"
// @Router /place/create [post]
func CreatePlace(c *gin.Context) {
	var place model.Place
	if err := c.BindJSON(&place); err != nil {
		handler.SendBadRequest(c, "Lack Param or Param Not Satisfiable.", nil, err)
		return
	}

	if err := model.DB.Create(&place).Error; err != nil {
		handler.SendError(c, "添加地点信息失败!", nil, err)
		return
	}

	handler.SendResponse(c, "添加地点信息成功!", nil)
}
