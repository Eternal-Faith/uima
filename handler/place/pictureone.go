package place

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

// @Summary "修改地点图片1"
// @Description "修改地点图片1"
// @Tags place
// @Accept application/json
// @Produce application/json
// @Param file formData file true "文件"
// @Param id formData string true "id--地点的id"
// @Success 200 {object} model.Script "{"mgs":"success"}"
// @Failure 400 "上传失败,请检查token与其他配置参数是否正确"
// @Router /place/pictureone [post]
func ModifyPlacePictureOne(c *gin.Context) {
	// c.Header("Access-Control-Allow-Origin", "*")
	file, err := c.FormFile("file")
	ID := c.PostForm("id") 
	
	log.Println("name is", ID)

	PATH := "places"

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
	scriptInfo, _ := model.GetPlaceInforbyId(ID)
	if scriptInfo.Path1 != "" && scriptInfo.Sha1 != "" {
		connector_github.RepoCreate().Del(scriptInfo.Path1, scriptInfo.Sha1)
	}

	// 上传新头像
	Base64 := services.ImagesToBase64(filename)
	picUrl, picPath, picSha := connector_github.RepoCreate().Push(PATH, file.Filename, Base64)
	fmt.Println(picUrl, picPath, picSha)
	os.Remove(filename)

	//这里大写是因为和上面重名了
	var Place model.Place
	Place.Id, _ = strconv.Atoi(ID)
	Place.Picture1 = picUrl
	Place.Path1 = picPath
	Place.Sha1 = picSha
	err = model.UpdatePlacePictureone(Place)
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