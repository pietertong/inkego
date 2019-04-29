package gojson

import (
    "encoding/json"
    "fmt"
)

type Student struct {
    Name string
    Age int
    Guake bool
    Classes []string
    Price float32
}

func (s * Student)ShowStu()  {
    
    fmt.Println("show Student : ")
    fmt.Println("\tName\t:",s.Name)
    fmt.Println("\tAge\t:",s.Age)
    fmt.Println("\tGuake\t:",s.Guake)
    fmt.Println("\tPrice\t:",s.Price)
    fmt.Println("\tClasses\t:")
    for _,a := range s.Classes{
        fmt.Printf("%s ",a)
    }
    fmt.Println("")
}

func EncodeJson()  {
    st := &Student{
        "Xiao Ming",
        16,
        true,
        [] string{"Math", "English", "Chinese"},
        9.99,
    }
    
    fmt.Println("Before Json Encoding : ")
    st.ShowStu()
    
    b,err := json.Marshal(st)
    if err != nil{
        fmt.Println("Encoding faild.")
    }else {
        fmt.Println("encoded data : ")
        //fmt.Println(b)
        fmt.Println(string(b))
    }
}

