package user

import (
	"strconv"
	"uima/handler"
	"uima/model"

	"github.com/gin-gonic/gin"
)

// @Summary "收藏剧本"
// @Description "将该剧本收藏到我的收藏中"
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param token header string true "token"
// @Param object body ScriptId true "剧本Id"
// @Success 200 "收藏成功"
// @Failure 401 "身份验证失败"
// @Failure 500 "服务器错误，收藏失败"
// @Router /user/collection [post]
func ScriptCollection(c *gin.Context) {
	users_id := c.MustGet("id").(int)

	var scriptId ScriptId
	if err := c.BindJSON(&scriptId); err != nil {
		handler.SendBadRequest(c, "Lack Param or Param Not Satisfiable.", nil, err)
		return
	}

	var collection model.ScriptCollections
	collection.UsersId = users_id
	collection.ScriptsId, _ = strconv.Atoi(scriptId.Id)

	if err := model.DB.Create(&collection).Error; err != nil {
		handler.SendError(c, "收藏失败", nil, err)
		return
	}

	handler.SendResponse(c, "收藏剧本成功！", nil)
}
