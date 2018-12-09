package chat_eight

import "fmt"

type IntVector[] int

func (v IntVector) Sum() (s int) {
    for _, x := range v {
        s += x
    }
    return
}

func eight_onetwo_main()  {
    fmt.Println(IntVector{1,2,3}.Sum())
}
