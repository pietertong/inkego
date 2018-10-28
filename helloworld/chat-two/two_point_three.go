package chat_two

import "fmt"

/**
2.3 类型
*/

/*
   布尔类型 ：bool
   整型 ：int8 , byte , int16 , int , uint , uintptr etc.
   浮点类型 ：float32 , float64
   复数类型 : complex64 , complex128
   字符串 ：string
   字符类型 ：rune
   错误类型 ：error
*/

/**
  指针(pointer)
  数组(array)
  切片(slice)
  字典(map)
  通道(chan)
  结构体(struct)
  接口(interface)
*/

/**
2.3.1
*/

/**
bool 值可赋值为预订的 true 和 false
不接受其他类型的赋值，不支持自动或强制的类型转换,
一下用法是正确的
*/
func main() {
	var b bool
	b = (1 != 0)
	fmt.Println("Result : ", b)
}

/**
int 与 int32 在GO语言里被认为是两种不同的类型，编译器也不会帮助你自动做类型转换
*/

/**
当然，开发者在做强制类型转换的时候，需要注意的是数据长度被截断而发生的数据精度损失和值溢出的问题
*/

/*
数值运算，
    var a float32
    a = -10 % -3
    fmt.Println("Result : ",a)
    //Result :  -1

    var a float32
    a = -10 % 3
    fmt.Println("Result : ",a)
    //Result :  -1

    var a float32
    a = 10 % -3
    fmt.Println("Result : ",a)
    //Result :  1
*/

/**
两个不同类型的整型数据不能直接比较，
比如 int8 类型的数据和 int 类型的数据不能直接比较，
但各种类型的整型变量都可以直接与字面常量进行比较
*/
func geteq() {
	var i int32
	var j int64
	i, j = 1, 2

	/* mismatched types int32 and int64
	   if i == j {

	       fmt.Println("Result : i and j equal.")
	   }
	*/
	if j == 2 || i == 1 {
		fmt.Println("Result : i eq 1 and j eq 2.")
	}
}

/**
2.3.3 浮点类型

浮点类型用于表示包含小数点的书， Go语言中的浮点类型采用IEEE-754标准的表达式
*/

/**
Go 语言中定义了 两个类型的 float32 、float64 ,分别等价于 C 语言中的float 、double
*/

func printLn() {
	var complex_value complex64
	complex_value = 12 + 3i

	fmt.Println("Result : real number = ", real(complex_value), " image number = ", imag(complex_value))
}

/**
字符串的内容可以用类似数组下标的方式获取，但与数组不同，字符串的内容不能再初始化后被修改


*/

/**
数组的声明




Go语言支持用 myArray[first:last]
这样的方式来基于数组生成一个数组切片，而且这个用法还很灵活
*/
