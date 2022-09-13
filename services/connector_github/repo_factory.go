package connector_github

import (
	"uima/services/flag_handle_fyc"
	"uima/services/github"
	"uima/services"
)

//定义serve的映射关系
var serveMap = map[string]services.RepoInterfacefyc{
	"github": &github.GithubServe{},
}

func RepoCreate()services.RepoInterfacefyc{
	return serveMap[flag_handle_fyc.PLATFORM]
}
