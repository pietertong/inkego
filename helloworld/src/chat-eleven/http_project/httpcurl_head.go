package http_project

import (
	//"encoding/json"
	"fmt"
	//"io/ioutil"
	"io"
	"log"
	"net/http"
	"os"
)

type RequestJsonUntil struct {
	author string `json:"author"`
}

/**
* http.Get() 相同
 */
func HttpCurlHeadMain() {
	fmt.Println("")
	requestUrl := "http://192.168.136.156:8081/"
	resp, err := http.Head(requestUrl)
	if err != nil {
		log.Fatal("Request " + requestUrl + " Error")
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
	//retBody, err := ioutil.ReadAll(resp.Body)
	//
	//fmt.Println(retBody)
	//if err != nil {
	//
	//}
	//var RequestJsonUntilAuthor RequestJsonUntil
	//
	//if err := json.Unmarshal(retBody, &RequestJsonUntilAuthor); err != nil {
	//
	//}
	//
	//fmt.Println("Request author of json : ",RequestJsonUntilAuthor.author)
}
