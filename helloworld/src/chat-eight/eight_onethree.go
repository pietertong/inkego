package chat_eight

import (
    "fmt"
    "time"
)

type myTime struct {
    time.Time
}

func (mt myTime) first3Chars() string {
    return mt.Time.String()[0:3]
}

func eight_onethree_main()  {
    m := myTime{time.Now()}
    //调用匿名 Time 上的String方法
    fmt.Println("完整的时间格式： ",m.String())
    //调用myTime.first3Chars
    fmt.Println("前三个字符： ",m.first3Chars())
}