package shop

import (
	"fmt"
	"os"
	"path"
	"time"
	"uima/handler"
	"uima/model"
	"uima/services"
	"uima/services/connector_github"

	"github.com/gin-gonic/gin"
)

// @Summary "店铺信息"
// @Tags shop
// @Description "新增店铺信息"
// @Accept json
// @Produce json
// @Param shop body model.Shop true "shop"
// @Success 200 "上传成功"
// @Failure 401 "身份验证失败"
// @Failure 400 "上传失败"
// @Router /shop [post]
func CreateShop(c *gin.Context) {

	var shop model.Shop
	file, err := c.FormFile("file")
	shop.CurrentNum = c.Request.FormValue("current_num")
	shop.FieryNum = c.Request.FormValue("fiery_num")
	shop.HotLine = c.Request.FormValue("hot_line")
	shop.OpeningTime = c.Request.FormValue("opening_time")
	shop.ServiceIntro = c.Request.FormValue("service_intro")
	shop.ShopName = c.Request.FormValue("shop_name")
	shop.ShouldKnow = c.Request.FormValue("should_know")
	shop.VipService = c.Request.FormValue("vip_service")
	// if err := c.BindJSON(&shop); err != nil {
	// 	handler.SendBadRequest(c, "Lack Param or Param Not Satisfiable.", nil, err)
	// 	return
	// }
	if err := model.DB.Table("shops").Create(&shop).Error; err != nil {
		handler.SendError(c, "添加商店信息失败", nil, err)
		return
	}
	if err != nil {
		fmt.Println(err)
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

	//	id1 := strconv.Itoa(user_id)
	//time.Now().Format("2006-01-02 15:04:00")
	//本来想加个时间戳防止格式不同时图床里出现两个文件，用户获取失败，不过这里好像删掉文件扩展名就行了
	//还是加个时间戳吧，扩展名还是有必要知道的
	//	file.Filename = user_id + time.Now().Format("2006-01-02 15:04:00") + fileExt
	file.Filename = time.Now().Format("20060102150400") + fileExt

	filename := filepath + file.Filename
	fmt.Println(filename)
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
	if shop.Path != "" && shop.Sha != "" {
		connector_github.RepoCreate().Del(shop.Path, shop.Sha)
	}
	PATH := "shop"
	// 上传新头像
	Base64 := services.ImagesToBase64(filename)
	picUrl, picPath, picSha := connector_github.RepoCreate().Push(PATH, file.Filename, Base64)

	os.Remove(filename)

	shop.Picture = picUrl
	var avatar model.Shop
	avatar.Id = shop.Id
	avatar.Picture = picUrl
	avatar.Path = picPath
	avatar.Sha = picSha

	err0 := model.UpdateShopAvator(avatar)
	if picUrl == "" || err0 != nil {
		c.JSON(401, gin.H{
			"message": "上传失败",
		})
		fmt.Println(3)
		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "创建成功",
		"shop_id": shop.Id,
		"url":     picUrl,
	})
}

// // @Summary "注册"
// // @Description "注册一个新用户"
// // @tags shop
// // @Accept json
// // @Produce json
// // @Param shop body model.User true "shop"
// // @Success 200 "用户创建成功"
// // @Failure 400 "输入有误，格式错误"
// // @Failure 401 "电话号码重复"
// // @Router /shop [post]
// func Register(c *gin.Context) {
// 	var shop model.User
// 	file, err := c.FormFile("file")
// 	shop.Phone = c.Request.FormValue("phone")
// 	shop.Password = c.Request.FormValue("password")
// 	shop.ConfirmPassword = c.Request.FormValue("confirm_password")

// 	//电话位数问题前端处理
// 	fmt.Println(shop.Phone)
// 	if shop.Password != shop.ConfirmPassword {
// 		c.JSON(400, gin.H{
// 			"code":    400,
// 			"message": "两次密码输入不一致，请重新输入"})
// 		return
// 	}
// 	if _, a := model.IfExistUserPhone(shop.Phone); a != 1 {
// 		c.JSON(401, gin.H{
// 			"code":    401,
// 			"message": "对不起，该电话号码已经被绑定",
// 		})
// 		return
// 	}
// 	shop.Password = model.GeneratePasswordHash(shop.Password)
// 	user_id := model.Register(shop.Phone, shop.Password)

// 	fmt.Println(shop.Phone, user_id)

// 	if err != nil {
// 		fmt.Println(err)
// 		fmt.Println(1)
// 		c.JSON(400, gin.H{
// 			"code":    401,
// 			"message": "上传失败!",
// 		})
// 		return
// 	}

// 	filepath := "./"
// 	if _, err := os.Stat(filepath); err != nil {
// 		if !os.IsExist(err) {
// 			os.MkdirAll(filepath, os.ModePerm)
// 		}
// 	}

// 	fileExt := path.Ext(filepath + file.Filename)

// 	//	id1 := strconv.Itoa(user_id)
// 	//time.Now().Format("2006-01-02 15:04:00")
// 	//本来想加个时间戳防止格式不同时图床里出现两个文件，用户获取失败，不过这里好像删掉文件扩展名就行了
// 	//还是加个时间戳吧，扩展名还是有必要知道的
// 	//	file.Filename = user_id + time.Now().Format("2006-01-02 15:04:00") + fileExt
// 	file.Filename = time.Now().Format("2006-01-02 15:04:00") + fileExt

// 	filename := filepath + file.Filename
// 	fmt.Println(filename)
// 	if err := c.SaveUploadedFile(file, filename); err != nil {
// 		fmt.Println(err)
// 		fmt.Println(2)
// 		c.JSON(400, gin.H{
// 			"code":    401,
// 			"message": "上传失败!",
// 		})
// 		return
// 	}

// 	// 删除原头像
// 	if shop.Path != "" && shop.Sha != "" {
// 		connector_github.RepoCreate().Del(shop.Path, shop.Sha)
// 	}

// 	// 上传新头像
// 	Base64 := services.ImagesToBase64(filename)
// 	picUrl, picPath, picSha := connector_github.RepoCreate().Push(file.Filename, Base64)

// 	os.Remove(filename)
// 	var avatar model.User
// 	avatar.Id = shop.Id
// 	avatar.Avatar = picUrl
// 	avatar.Path = picPath
// 	avatar.Sha = picSha
// 	err0 := model.UpdateAvator(avatar)
// 	if picUrl == "" || err0 != nil {
// 		c.JSON(401, gin.H{
// 			"message": "上传失败",
// 		})
// 		fmt.Println(3)
// 		return
// 	}

// 	c.JSON(200, gin.H{
// 		"code":    200,
// 		"message": "上传成功",
// 		"url":     picUrl,
// 		"sha":     picSha,
// 		"path":    picPath,
// 	})

// 	fmt.Println(shop.Id)
// 	shop.Id, _ = strconv.Atoi(user_id)
// 	signedToken := model.GenerateToken(shop.Id)
// 	c.JSON(200, gin.H{
// 		"code":    200,
// 		"message": "将用户id作为token保留",
// 		"data":    signedToken,
// 	})
// }
