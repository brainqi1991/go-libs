package web

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	// "strings"
	"bytes"
	"io/ioutil"
)

// Get 发送 get 请求给solr服务器，返回 body中的字符串内容&状态码，方便解析转换成JSON
func Get(url string) (string, int) {
	client := &http.Client{}
	// var url string = "http://localhost:9091/user/login/1"
	// url = "http://test.com:8983/solr/rank-queen/select?q=*:*&fq=id:198784&wt=json&indent=true"
	response, _ := client.Get(url)
	defer response.Body.Close()
	// log.Printf("返回信息 cookies: %v\n request: %v\n body: %v\n headers:%v\n\n\nStatusCode: %v\n",
	// 	response.Cookies(), response.Request, response.Body, response.Header, response.StatusCode)
	result, _ := ioutil.ReadAll(response.Body)
	log.Println(string(result))
	return string(result), response.StatusCode
}

// PostForm 发送 post 请求参数封装成form
// 给solr服务器，返回 body中的字符串内容&状态码，方便解析转换成JSON
func PostForm(url string, data url.Values) (string, int) {
	client := &http.Client{}
	response, _ := client.PostForm(url, data)
	defer response.Body.Close()
	result, _ := ioutil.ReadAll(response.Body)
	log.Println(result)
	return string(result), response.StatusCode
}

// PostJson 支持 JSON字符串和 struct 参数发送请求给服务器
func PostJson(url string, v interface{}) (string, int) {
	client := &http.Client{}
	var str []byte
	if value, ok := v.(string); ok {
		str = []byte(value)
	} else {
		str, _ = json.Marshal(v)
	}
	data := bytes.NewReader(str)
	response, _ := client.Post(url, "application/json", data)
	defer response.Body.Close()
	result, _ := ioutil.ReadAll(response.Body)
	log.Println(string(result))
	return string(result), response.StatusCode
}
