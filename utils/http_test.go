package utils

import (
	"go-awesome-webapp/model"
	"net/url"
	"testing"
)

func Test_Get(t *testing.T) {
	// var url string = "http://localhost:9091/user/login/1"
	var url = "http://test.com:8983/solr/rank-queen/select?q=*:*&fq=id:198784&wt=json&indent=true"
	result, code := Get(url)

	if code != 200 {
		t.Errorf("服务器响应不正确，StatusCode : %d", code)
	} else {
		t.Logf("测试通过，Result : %s", result)
	}
}

func Test_PostForm(t *testing.T) {

	urlStr := "http://localhost:9091/lianxi-queen/web/activity_dynamic/dynamic_get"
	form := make(url.Values)
	form.Set("name", "brainqi@yeah.net")
	form.Set("secret", "giheum")
	result, code := PostForm(urlStr, form)

	if code != 200 {
		t.Errorf("服务器响应不正确，StatusCode : %d", code)
	} else {
		t.Logf("测试通过，Result : %s", result)
	}
}

func Test_PostJson(t *testing.T) {
	// json := `{"name":"BrainQi"}`
	json := model.UserJsonForm{Name: "BrainQi", Secret: "123"}
	urlStr := "http://localhost:9091/lianxi-queen/web/activity_dynamic/dynamic_get"
	result, code := PostJson(urlStr, json)
	if code != 200 {
		t.Error("JSON发送测试失败!")
	} else {
		t.Logf("嗯哼，Result: %v", result)
	}
}
