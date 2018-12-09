package chat_eight

import (
    "fmt"
    "reflect"
)

/**
GoLang 获取struct 的tag
 */

type CEuser struct {
    Name string "user name"
    Passwd string "user password"
}

func eight_struct_tags() {
    user := &CEuser{"chronos","pass"}
    s := reflect.TypeOf(user).Elem()
    
    for i := 0; i < s.NumField(); i++ {
        fmt.Println(s.Field(i).Tag)
    }
}

