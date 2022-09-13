package main

import (
	"flag"
	"fmt"
	"uima/config"
	"uima/model"
	"uima/router"

	//"uima/services/flag_handle"
	"uima/services/flag_handle_fyc"

	"github.com/gin-gonic/gin"
	_ "github.com/spf13/viper"
	_ "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title uima
// @version 1.0.0
// @description 山乡剧变
// @termsOfService http://swagger.io/terrms
// @contact.name Eternal-Faith
// @contact.email 2295616516@qq.com
// @host localhost:8918
// @BasePath:/api/v1
// @Schemes http
func main() {
	r := gin.Default() //创建带有默认中间件的路由
	config.ConfigInit()
	//注意大写规范
	model.DB = model.Initdb()
	router.Router(r)
	if err := r.Run(":8795"); err != nil {
		fmt.Println(err)
	}

}

func init() {
	port := flag.String("port", "8795", "本地监听的端口")
	platform := flag.String("platform", "github", "平台名称，支持gitee/github")
	token := flag.String("token", "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", "Gitee/Github 的用户授权码")
	owner := flag.String("owner", "SUIYUELIANYI", "仓库所属空间地址(企业、组织或个人的地址path)")
	repo := flag.String("repo", "project_image", "仓库路径(path)")
	branch := flag.String("branch", "main", "分支")
	flag.Parse()
	flag_handle_fyc.PORT = *port
	flag_handle_fyc.OWNER = *owner
	flag_handle_fyc.REPO = *repo
	flag_handle_fyc.TOKEN = *token
	flag_handle_fyc.PLATFORM = *platform
	flag_handle_fyc.BRANCH = *branch
	if flag_handle_fyc.TOKEN == "" {
		panic("token 必须！")
	}
}
