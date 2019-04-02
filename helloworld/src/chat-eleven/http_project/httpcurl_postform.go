package http_project

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

/**
{
    "get": [],
    "post": {
        "id": "123",
        "key": "value"
    }
}
*/
func HttpCurlPostform() {
	fmt.Println("")
	requestUrl := "http://192.168.136.156:8081/"
	resp, err := http.PostForm(requestUrl, url.Values{"key": {"value"}, "id": {"123"}})
	if err != nil {
		log.Fatal("Request " + requestUrl + " Error")
	}
	defer resp.Body.Close()
	//把请求到的数据打印到标准输出流中
	io.Copy(os.Stdout, resp.Body)
}
