package rpc_project

import (
    "fmt"
    "log"
    "net/rpc"
    "os"
)

type ArgsClient struct {
    A,B int
}

type QuotientClient struct {
    Quo, Rem int
}

func GorpcClient() {
    if len(os.Args) != 2{
        fmt.Println("用法 ： ", os.Args[0],"server")
        os.Exit(1)
    }
    
    serverAddress := os.Args[1]
    
    client, err := rpc.DialHTTP("tcp", serverAddress + ":1234")
    if err != nil{
        log.Fatal("Dial 错误 ： ", err)
    }
    
    args := ArgsClient{17, 8}
    var reply int
    
    err = client.Call("Arith.Multiply", args, &reply)
    if err != nil {
        log.Fatal("Arith 错误 ： ", err)
    }
    fmt.Printf("Arith : %d * %d = %d ./n", args.A, args.B, reply)
    
    var quot QuotientClient
    err = client.Call("Arith.Divide", args, &quot)
    if err != nil {
        log.Fatal("Arith 错误 ： ", err)
    }
    fmt.Printf("Arith : %d / %d = %d 余 %d./n", args.A, args.B, quot.Quo, quot.Rem)
    
}
