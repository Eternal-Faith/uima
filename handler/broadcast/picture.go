package broadcast

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"time"
	"uima/model"
	"uima/services"
	"uima/services/connector_github"

	"github.com/gin-gonic/gin"
)

// @Summary "修改精彩放送图片"
// @Tags bro
// @Description "精彩放送图片"
// @Accept json
// @Produce json
// @Param broadcast_id path string true "broadcast_id"
// @Param file formData file true "文件"
// @Success 200 "上传成功"
// @Failure 401 "身份验证失败"
// @Failure 400 "上传失败"
// @Router /broadcast/avatar/:broadcast_id [post]
func ModifyBroProfile(c *gin.Context) {
	id := c.Param("broadcast_id")
	// var broadcast model.Broadcast

	bro, _ := model.GetSingleBroadcastInfo(id)
	file, err := c.FormFile("file")

	if err != nil {
		fmt.Println(1)
		c.JSON(400, gin.H{
			"code":    401,
			"message": "上传失败!",
		})
		return
	}

	filepath := "./"
	if _, err := os.Stat(filepath); err != nil {
		if !os.IsExist(err) {
			os.MkdirAll(filepath, os.ModePerm)
		}
	}

	fileExt := path.Ext(filepath + file.Filename)

	//time.Now().Format("2006-01-02 15:04:00")//time这个格式打开文件会有文件名或路径的报错，把空格斜杠特殊符号删掉即可
	//本来想加个时间戳防止格式不同时图床里出现两个文件，用户获取失败，不过这里好像删掉文件扩展名就行了
	//还是加个时间戳吧，扩展名还是有必要知道的
	file.Filename = "bro" + time.Now().Format("20060102150304") + fileExt

	filename := filepath + file.Filename
	if err := c.SaveUploadedFile(file, filename); err != nil {
		fmt.Println(err)
		fmt.Println(2)
		c.JSON(400, gin.H{
			"code":    401,
			"message": "上传失败!",
		})
		return
	}

	// 删除原头像

	if bro.Path != "" && bro.Sha != "" {
		connector_github.RepoCreate().Del(bro.Path, bro.Sha)
	}

	PATH := "bro"
	// 上传新头像
	Base64 := services.ImagesToBase64(filename)
	picUrl, picPath, picSha := connector_github.RepoCreate().Push(PATH, file.Filename, Base64)

	os.Remove(filename)
	var avatar model.Broadcast
	Id, _ := strconv.Atoi(id)
	avatar.Id = Id
	avatar.Picture = picUrl
	avatar.Path = picPath
	avatar.Sha = picSha
	err0 := model.UpdateBroAvator(avatar)
	fmt.Println(avatar)
	if picUrl == "" || err0 != nil {
		c.JSON(401, gin.H{
			"message": "上传失败",
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "上传成功",
		"url":     picUrl,
		"sha":     picSha,
		"path":    picPath,
	})

}
