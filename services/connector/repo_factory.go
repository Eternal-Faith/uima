package connector

import (
	"uima/services"
	"uima/services/flag_handle"
	"uima/services/gitee"
)

//定义serve的映射关系
var serveMap = map[string]services.RepoInterface{
	"gitee": &gitee.GiteeServe{},
}

func RepoCreate() services.RepoInterface {
	return serveMap[flag_handle.PLATFORM1]
}
