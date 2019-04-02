package http_project

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func HttpCurlPostMain() {
	fmt.Println("")
	postData := "{\"protect\": {\"go\": [\"post_json_data\"],\"post_json_data\": {\"l1_2_1\": 121}}}"
	var jsonStr = []byte(postData)
	fmt.Println("jsonStr", jsonStr)
	newStr := bytes.NewBuffer(jsonStr)
	fmt.Println("new_str", newStr)
	requestUrl := "http://192.168.136.156:8081/"
	resp, err := http.Post(requestUrl, "application/json;charset=utf-8", newStr)

	if err != nil {
		log.Fatal("Request " + requestUrl + " Error")
	}
	defer resp.Body.Close()
	//把请求到的数据打印到标准输出流中
	io.Copy(os.Stdout, resp.Body)

}
