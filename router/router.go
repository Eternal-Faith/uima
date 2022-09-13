package router

import (
	"net/http"
	"uima/handler/auth"
	"uima/handler/broadcast"
	"uima/handler/place"
	"uima/handler/script"
	"uima/handler/shop"
	"uima/handler/user"
	"uima/router/middleware"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API router.")
	})

	//auth:
	g1 := r.Group("/api/v1/auth")
	{
		g1.POST("/register", auth.Register)

		g1.POST("/login", auth.Login)
	}

	//user:
	g2 := r.Group("/api/v1/user").Use(middleware.Auth())
	{
		//修改个人信息
		g2.POST("/edit", user.Edit)

		//修改头像
		g2.POST("/avatar", user.ModifyProfile)

		//游览个人信息
		g2.POST("/view", user.Userinfo)

		//我的收藏(剧本)
		g2.GET("/mycollection", user.MyCollection)

		//查看订单(剧本)
		g2.GET("/vieworder", user.ViewOrder)

		//查看预约(剧本)
		g2.GET("/myappointment", user.MyAppointment)

		//根据状态查看预约(剧本)
		g2.POST("/searchappointment",user.SearchAppointment)

		//订单付款(剧本)
		g2.PUT("/payforscript",user.PayforScript)

		//更新预约状态
		g2.PUT("/updateappointmentstatus",user.UpdateAppointmentStatus)

		//取消预约
		g2.POST("/cancelappoint",user.CancelScriptAppoint)

		//收藏剧本
		g2.POST("/collection", user.ScriptCollection)

		//取消收藏
		g2.POST("/cancel", user.CancelScriptCollection)

		//预约剧本
		g2.POST("/appointment", user.Appointment)

		//创建订单
		g2.POST("/order", user.CreateScriptOrder)
	}

	//script:
	g3 := r.Group("/api/v1/script")
	{
		//创建剧本
		g3.POST("/create", script.CreateScript)

		//修改剧本封面
		g3.POST("/cover", script.ModifyScriptCover)

		//获取剧本主界面
		g3.GET("/interface", script.Interface)

		//查看剧本
		g3.POST("/view", script.ViewScript)

		//标签搜索
		g3.POST("/tag", script.TagSearch)

		//修改剧本
		g3.POST("/edit", script.EditScript)

	}

	//place:
	g4 := r.Group("/api/v1/place")
	{
		//增加地点信息
		g4.POST("/create", place.CreatePlace)

		//查询地点信息
		g4.POST("/view", place.ViewPlace)

		//修改地点图片1
		g4.POST("/pictureone", place.ModifyPlacePictureOne)

		//修改地点图片2
		g4.POST("/picturetwo", place.ModifyPlacePictureTwo)

		//修改地点图片3
		g4.POST("/picturethree", place.ModifyPlacePictureThree)
	}

	//broadcast:
	g5 := r.Group("/api/v1/broadcast")
	{
		//增加精彩放送信息
		g5.POST("/basic_info", broadcast.CreateBroadcast)

		//设置精彩放送图片
		g5.POST("/avatar/:broadcast_id", broadcast.ModifyBroProfile)

		//获取所有精彩放送信息
		g5.GET("/get_all", broadcast.GetBroadcast)

		//获取单个精彩放送信息
		g5.GET("/get_single/:broadcast_id", broadcast.GetSingleBroadcast)

		//获取单个精彩放送信息
		g5.DELETE("/:broadcast_id", broadcast.DeleteBroadcast)

	}

	//shop:
	g6 := r.Group("/api/v1/shop")
	{
		//增加商店信息
		g6.POST("", shop.CreateShop)

		//获取所有商店信息
		g6.GET("/get_all", shop.GetAllShopInfo)

		//获取单个商店信息
		g6.GET("/single_shop/:shop_id", shop.GetSingleShopInfo)

		//获取单个精彩放送信息
		g6.DELETE("/:shop_id", shop.DeleteShop)
	}
}
