package chat_five

import (
	"common"
	"fmt"
)

/**
函数的首字母大写表示 可以在最外层的package 调用
首字母小写表示 当前文件内或当前的package 可以调用
*/
func ChatFive() {
	var formatString common.Funclib

	formatString = "%d,%d"

	fmt.Println(formatString)

	s1 := pipe(func() int {
		return 100
	})
	s2 := format(func(s string, x, y int) string {
		return fmt.Sprintf(s, x, y)
	}, "%d,%d", 102, 103)

	fmt.Println(fmt.Sprintf("sprint ln : %d,%d", 10, 20))
	fmt.Println(s1, s2)

	/**
	首字母小写调用，当前包的函数方法
	*/
	helloworld()
	/**
	调用闭包函数
	 */
	fmt.Printf("%d." , add()(3))
	fmt.Println("")
	fmt.Printf("%d." , add2(6)(3))


	/**
	调用闭包函数，继承当前函数内的作用域问题
	 */
	 closurfunc()
}

func pipe(ff func() int) int {
	return ff()
}

type FormatFunc func(s string, x, y int) string

func format(ff FormatFunc, s string, x, y int) string {
	return ff(s, x, y)
}
