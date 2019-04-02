package http_diy_project

import (
	"log"
	"net/http"
)

//func redirectPolicyFunc()  {
//    log.Fatal("redirect Policy Func")
//}
func DiyClientFuncMain() {
	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}
	requestUrl := "http://192.168.136.156:8081/"
	resp, err := client.Get(requestUrl)
	if err != nil {
		log.Fatal("Request " + requestUrl + " error")
	}
	defer resp.Body.Close()

}
