package chat_eight

import (
    "fmt"
    "strings"
)
type Person struct {
    age int8
    firstName string
    lastName string
}

func (pointer *Person) echo(age int8)(Person ,error) {
    fmt.Println("This is echo func .")
    var temp  Person
    temp.age = age
    return temp,nil
}

func upPerson(p *Person)  {
    p.firstName = strings.ToUpper(p.firstName)
    p.lastName = strings.ToUpper(p.lastName)
}

func eight_two_main()  {
    var pers1 Person
    pers1.firstName = "张"
    pers1.lastName = "三三"
    
    upPerson(&pers1)
    var temp,_ = pers1.echo(10)
    fmt.Println(temp)
    fmt.Printf("这个人的姓名是 %s %s .\n",pers1.firstName,pers1.lastName)
}