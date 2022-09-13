package script

import (
	"uima/handler"
	"uima/model"

	"github.com/gin-gonic/gin"
)

type Scripts struct {
	ScriptName   string `gorm:"column:script_name;type:varchar(255);NOT NULL" json:"script_name"`    // 剧本名称
	Place        string `gorm:"column:place;type:varchar(255);NOT NULL" json:"place"`                // 地点
	Time         string `gorm:"column:time;type:varchar(255);NOT NULL" json:"time"`                  // 剧本时间
	Introduction string `gorm:"column:introduction;type:varchar(1000);NOT NULL" json:"introduction"` // 剧本介绍
	BriefIntro   string `gorm:"column:brief_intro;type:varchar(255);NOT NULL" json:"brief_intro"`    // 剧本简介
	Price        int    `gorm:"column:price;type:int(11);NOT NULL" json:"price"`                     // 价格
	Tag1         string `gorm:"column:tag1;type:varchar(100)" json:"tag1"`                           // 标签一
	Tag2         string `gorm:"column:tag2;type:varchar(100)" json:"tag2"`                           // 标签二
	Tag3         string `gorm:"column:tag3;type:varchar(100)" json:"tag3"`                           // 标签三
	Tag4         string `gorm:"column:tag4;type:varchar(100)" json:"tag4"`                           // 标签四
	Tag5         string `gorm:"column:tag5;type:varchar(100)" json:"tag5"`                           // 标签五
}

// @Summary "创建剧本"
// @Describtion "上传剧本的名称，地点，封面，介绍"
// @Tags script
// @Accept json
// @Producer json
// @Param object body Scripts true "剧本"
// @Success 200 "创建成功"
// @Failure 400 "输入格式有误"
// @Failure 500 "服务器错误"
// @Router /script/create [post]
func CreateScript(c *gin.Context) {
	var script model.Script
	if err := c.BindJSON(&script); err != nil {
		handler.SendBadRequest(c, "Lack Param or Param Not Satisfiable.", nil, err)
		return
	}

	if err := model.DB.Create(&script).Error; err != nil {
		handler.SendError(c, "创建剧本失败!", nil, err)
		return
	}

	handler.SendResponse(c, "创建剧本成功!", nil)
}
