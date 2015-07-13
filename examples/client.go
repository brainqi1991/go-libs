package main

import (
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{}

	var url string = "http://localhost:9091/user/login/1"
	url = "http://test.com:8983/solr/rank-queen/select?q=*:*&fq=id:198784&wt=json&indent=true"
	response, _ := client.Get(url)
	defer response.Body.Close()
	fmt.Printf("返回信息 cookies: %v\n request: %v\n body: %v\n headers:%v\n\n\nStatusCode: %v\n",
		response.Cookies(), response.Request, response.Body, response.Header, response.StatusCode)

	buffer := make([]byte, 1024)
	var result string
	for {
		n, _ := response.Body.Read(buffer)
		if n == 0 {
			break
		}
		result = string(buffer)
		// fmt.Printf("Result : %v", result)
	}
	fmt.Println("===================== 分割线 ==================")
	fmt.Println(result)
}
