package chat_eight

import "fmt"

type myStruct struct {
	i1  int
	f1  float32
	str string
}

func eight_main() {
	ms := new(myStruct)
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str = "Google"

	fmt.Println("Chat eight :")
	fmt.Printf("int : %d \n", ms.i1)
	fmt.Printf("float : %f \n", ms.f1)
	fmt.Printf("string : %s \n", ms.str)
	fmt.Println(ms)
}
