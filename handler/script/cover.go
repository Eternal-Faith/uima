package script

import (
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"uima/handler"
	"uima/model"
	"uima/services"
	"uima/services/connector_github"

	"github.com/gin-gonic/gin"
)

// @Summary "修改剧本封面"
// @Description "修改剧本封面"
// @Tags script
// @Accept application/json
// @Produce application/json
// @Param file formData file true "文件"
// @Param id formData string true "id--剧本的id"
// @Success 200 {object} model.Script "{"mgs":"success"}"
// @Failure 400 "上传失败,请检查token与其他配置参数是否正确"
// @Router /script/cover [post]
func ModifyScriptCover(c *gin.Context) {
	//c.Header("Access-Control-Allow-Origin", "*")
	file, err := c.FormFile("file")
	ID := c.PostForm("id") 
	
	log.Println("name is", ID)

	PATH := "scripts"

	if err != nil {
		handler.SendBadRequest(c, "上传失败!", nil, err)
		return
	} 

	filepath := "./"
	if _, err := os.Stat(filepath); err != nil {
		if !os.IsExist(err) {
			os.MkdirAll(filepath, os.ModePerm)
		}
	}

	fileExt := path.Ext(filepath + file.Filename)

	file.Filename = ID + "_" + services.GetRandomString(16) + fileExt

	filename := filepath + file.Filename

	if err := c.SaveUploadedFile(file, filename); err != nil {
		handler.SendBadRequest(c, "上传失败!!!", nil, err)
		return
	}

	// 删除原头像
	scriptInfo, _ := model.GetScriptInfor(ID)
	if scriptInfo.Path != "" && scriptInfo.Sha != "" {
		connector_github.RepoCreate().Del(scriptInfo.Path, scriptInfo.Sha)
	}

	// 上传新头像
	Base64 := services.ImagesToBase64(filename)
	picUrl, picPath, picSha := connector_github.RepoCreate().Push(PATH, file.Filename, Base64)
	fmt.Println(picUrl, picPath, picSha)
	os.Remove(filename)

	//这里大写是因为和上面重名了
	var Script model.Script
	Script.Id, _ = strconv.Atoi(ID)
	Script.Avatar = picUrl
	Script.Path = picPath
	Script.Sha = picSha
	err = model.UpdateScriptCover(Script)
	if picUrl == "" || err != nil {
		handler.SendBadRequest(c, "上传失败,请检查token与其他配置参数是否正确", nil, err)
		return
	}

	handler.SendResponse(c, "上传成功", map[string]interface{}{
		"url":  picUrl,
		"sha":  picSha,
		"path": picPath,
	})
}
