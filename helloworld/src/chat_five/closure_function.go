package chat_five

import "fmt"

/**
闭包
 */

func add() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}


func add2(a int) func(b int) int{
	return func(b int) int {
		return a + b
	}
}

/**
理解闭包继承函数声明是的作用域特定
 */
func closurfunc()  {
	j := 5
	a := func() func(){
		i := 10
		return func() {
			fmt.Printf("i = %d, j = %d.\n", i, j)
		}
	}()

	a()

	j = 10

	a()
}