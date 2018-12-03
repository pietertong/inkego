package chat_five

import "fmt"

/**
参数传递机制

1、按值传递
2、地址传递
*/

func advancedAdd(a int) int {
	a++
	return a
}

func advancedAddPoint(a *int) int {
	*a++
	return *a
}

func advanceFunc() {
	x := 3
	fmt.Println("x = ", x, "&x = ", &x)

	/**
	  地址传递
	*/
	y := advancedAdd(x)
	fmt.Println("x = ", x, "&y = ", &y)

	/**
	  值传递
	*/
	z := advancedAddPoint(&x)
	fmt.Println("x = ", x, "&x = ", &x)
	fmt.Println("z = ", z, "&z = ", &z)

}
