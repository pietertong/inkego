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

func add2(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

/**
理解闭包继承函数声明是的作用域特定
*/
func closurfunc() {
	j := 5
	a := func() func() {
		i := 10
		return func() {
			fmt.Printf("i = %d, j = %d.\n", i, j)
		}
	}()

	a()

	j = 10

	a()
}

func closurfunc_scope() {
	var f = adder()
	//相当于 fmt.Println(adder()(1))
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
}

/**

 */

func adder() func(int) int {
	var x int // 闭包中的变量可以在闭包函数体内申明，也可以在外部函数声明
	return func(d int) int {
		fmt.Println("x = ", x, "d = ", d)
		x += d
		return x
	}
}

func getSequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
