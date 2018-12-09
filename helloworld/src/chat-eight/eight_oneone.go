package chat_eight

import "fmt"

type TwoInts struct {
    a int
    b int
}

func eight_oneone_main()  {
    two1 := new(TwoInts)
    two1.a = 12
    two1.b = 10
    
    fmt.Println("a,b 的和为: %d;\n",two1.add() )
    
    //point_a := two1.add()
    //fmt.Printf("pointer address : %p . %p .",&TwoInts{},&point_a)
    
    fmt.Println("将他们添加到参数中:%d.\n",two1.addParams(20))
    
    two2 := TwoInts{3,4}
    
    fmt.Printf("和 为 ： %d.\n",two2.add())
}

func echo (ti *TwoInts)  {
    fmt.Printf("%p", ti)
}

func (ti *TwoInts) add() int {
    fmt.Printf("%d\n", ti.b + ti.b)
    return 0
}

func (ti *TwoInts) addParams(parasm int) int {
    return (ti.b + ti.b + parasm)
}
