package chat_eight

import "fmt"

/**
为类型添加方法
 */
 
type Integer int

func (a Integer) Less(b Integer) bool {
    return a < b
}

func eight_demo_main()  {
    var a Integer = 1
    if a.Less(2) {
        fmt.Println(a , "Less 2")
    } else {
        fmt.Println(a, "Else less 2")
    }
}