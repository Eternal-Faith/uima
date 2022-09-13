package shop

import (
	"uima/model"

	"github.com/gin-gonic/gin"
)

// @Summary "删除店铺信息"
// @Tags shop
// @Description "删除一个店铺的信息"
// @Accept json
// @Produce json
// @Param shop_id path string true "shop_id"
// @Success 200 "删除成功"
// @Failure 400 "删除失败"
// @Router /shop/:shop_id [delete]
func DeleteShop(c *gin.Context) {
	id := c.Param("shop_id")
	if err := model.DeleteSingleShopInfo(id); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "删除失败",
		})
		return
	}
	
	c.JSON(200, gin.H{
		"code":    200,
		"message": "删除成功",
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
// 		connector.RepoCreate().Del(shop.Path, shop.Sha)
// 	}

// 	// 上传新头像
// 	Base64 := services.ImagesToBase64(filename)
// 	picUrl, picPath, picSha := connector.RepoCreate().Push(file.Filename, Base64)

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
