package http_project

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func HttpCurlGetMain() {
	fmt.Println("")
	requestUrl := "http://192.168.136.156:8081/"
	resp, err := http.Get(requestUrl)
	if err != nil {
		log.Fatal("Request " + requestUrl + " Error")
	}
	defer resp.Body.Close()
	//把请求到的数据打印到标准输出流中
	io.Copy(os.Stdout, resp.Body)

	retBody, err := ioutil.ReadAll(resp.Body)

	fmt.Println(retBody)
}
