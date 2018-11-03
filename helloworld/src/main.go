package main

import (
	"fmt"
	"os"
	"common"
)

var pietertong common.Funclib

func main()  {
	pietertong = "zhaotong"
	fmt.Println(os.Getenv("GOPATH"))
	fmt.Printf("%s,%d",pietertong,30)
}
