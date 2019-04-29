package goweb

import (
    "common"
    "fmt"
    "log"
    "net/http"
    "strconv"
    "text/template"
)

func ConstructServices()  {
    http.HandleFunc("/",templateHandler)
    http.HandleFunc("/login",loginHandler)
    err := http.ListenAndServe(":9090",nil)
    if err !=  nil{
        log.Fatal("ListenAndServer:",err)
    }
}
func loginHandler(w http.ResponseWriter, r *http.Request) {
    method := r.Method
    fmt.Println("method : " ,method)
    if method == "GET" {
        templateFile := fmt.Sprintf("%s/%s",common.GetCurPath(),"templates/login.html")
        if common.Exists(templateFile) {
            t, _ := template.ParseFiles(templateFile)
            t.Execute(w, nil)
        }else{
            fmt.Println(fmt.Sprintf("%s not exist.",templateFile))
        }
    }else {
        r.ParseForm()
        username := r.Form["username"]
        password := r.Form["password"]
        fmt.Println("password:",password)
        fmt.Println("username:",username)
        if len(username) > 0 {
            fmt.Println("username's length more than : " ,0)
        }
        num := r.Form.Get("num")
        getint,err := strconv.Atoi(num)
        fmt.Print(fmt.Sprintf("%s,%s",num,getint))
        if err != nil{
            fmt.Println(fmt.Sprintf("%s","数字转化出错,那么可能就不是数字"))
        }
    }
}
func templateHandler(w http.ResponseWriter, r *http.Request){
    templateFile := fmt.Sprintf("%s/%s",common.GetCurPath(),"templates/layout.html")
    t, err :=template.ParseFiles(templateFile)
    if err != nil{
        log.Fatal("error",err)
    }
    fmt.Println(t.Name())
    t.Execute(w, "Hello world")
}
