package chat_five

/**
递归 函数 就是在运行的过程中，自己调用自己，需要递归条件满足之后，终止程序
*/

/**
阶乘递归
*/
func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

/**
斐波那契数列
*/

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}
