package chat_three

import "fmt"

/**
  第三章 面向对象编程 demo start
*/
func ChatThree() {
	demo_integer()
}

/**
  第三章 面向对象编程 demo end
*/

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

func demo_integer() {
	var a Integer = 1
	if a.Less(2) {
		fmt.Println(a, "Less 2")
	}
}
