package goenvsystem

import (
	"fmt"
	"os"
)

func GetSystemGoenv() {
	fmt.Printf("项目的 GORoot : %s .\n", os.Getenv("GOROOT"))
	fmt.Printf("项目的 GOPath : %s .\n", os.Getenv("GOPATH"))
}
