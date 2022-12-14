package gitee

import (
	"encoding/json"
	"fmt"
	"log"
	"uima/services"
	"uima/services/flag_handle"

	"github.com/valyala/fasthttp"
)

type GiteeServe struct {
	services.RepoInterface
}

func (serve *GiteeServe) Push(filename, content string) (string, string, string) {
	return PushGitee(filename, content)
}

func (serve *GiteeServe) Del(filepath, sha string) string {
	return DelFile(filepath, sha)
}

func PushGitee(filename, content string) (string, string, string) {
	url := "https://gitee.com/api/v5/repos/" + flag_handle.OWNER1 + "/" + flag_handle.REPO1 + "/contents/" + flag_handle.PATH1 + "/" + filename

	// 初始化请求与响应
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		// 用完需要释放资源
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()

	// 设置请求方法
	req.Header.SetMethod("POST")
	req.Header.SetBytesKV([]byte("Content-Type"), []byte("application/json"))

	// 设置请求的目标网址
	req.SetRequestURI(url)

	args := make(map[string]string)
	args["access_token"] = flag_handle.TOKEN1
	args["content"] = content
	args["message"] = "upload pic for blackboard"
	args["branch"] = flag_handle.BRANCH1

	jsonBytes, _ := json.Marshal(args)
	req.SetBodyRaw(jsonBytes)

	// 发送请求
	if err := fasthttp.Do(req, resp); err != nil {
		log.Println(" 请求失败:", url, err.Error())
	}

	// 获取相应的数据实体
	body := resp.Body()

	var mapResult map[string]interface{}

	err := json.Unmarshal(body, &mapResult)
	if err != nil {
		fmt.Println("JsonToMapDemo err :", err)
	}

	d := ""
	p := ""
	s := ""
	_, ok := mapResult["content"]

	if ok {
		if mapResult["content"] != nil {
			d = mapResult["content"].(map[string]interface{})["download_url"].(string)
			p = mapResult["content"].(map[string]interface{})["path"].(string)
			s = mapResult["content"].(map[string]interface{})["sha"].(string)
		}
	}
	return d, p, s
}

func DelFile(filepath, sha string) string {

	url := "https://gitee.com/api/v5/repos/" + flag_handle.OWNER1 + "/" + flag_handle.REPO1 + "/contents/" + filepath

	// 初始化请求与响应
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	defer func() {
		// 用完需要释放资源
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()

	// 设置请求方法
	req.Header.SetMethod("DELETE")
	req.Header.SetBytesKV([]byte("Content-Type"), []byte("application/json"))

	// 设置请求的目标网址
	req.SetRequestURI(url)

	args := make(map[string]string)
	args["access_token"] = flag_handle.TOKEN1
	args["sha"] = sha
	args["message"] = "delete pic for repo-image-hosting"
	args["branch"] = flag_handle.BRANCH1

	jsonBytes, _ := json.Marshal(args)
	req.SetBodyRaw(jsonBytes)

	// 发起请求
	if err := fasthttp.Do(req, resp); err != nil {
		log.Println(" 请求失败:", url, err.Error())
	}

	// 获取响应的数据实体
	body := resp.Body()

	return string(body)
}
