/**
* 注释风格与 C/C++ 保持一致，GO不要钱每个语句后边都加入分号,
* 左花括号 不能另起一行，否则会报编译错误 syntax error : unexpected semiconlon or newline before {
 */

//该代码所属的包名为 main;包是Go语言里的最基本的分发单位，也是工程管理中的依赖关系的体现
//要生产 GO可执行的程序，必须建立一个名为 main 的包，并且需要在包中包含一个叫main()的函数，是执行程序的执行起点
package main

//导入改程序需要的依赖包，不要在源代码中未使用的包，否则会报编译错误
import "fmt"

// main()函数不能带任何参数，也不能定义返回值，
//命令行传入的参数是在 os.Args 变量中保存

func main() {
	fmt.Println("Go,Lang,Hello,World!")
	_, _, last_name := get_name()
	fmt.Println(last_name)
}

/**
匿名函数
*/
func get_name() (frist_name, second_name, last_name string) {
	return "Zhao", "shu", "yu"
}

/**
func 函数名(参数列表)(返回值列表){
    函数体
}
*/

/**
  第二章 顺序编程-基础 demo start
*/
func chat_two()  {

	demo_boolean()
	demo_int_equal()
	demo_complex()
	demo_string()
	demo_foreach()
	demo_slice()
	demo_slice_two()
	demo_slice_two_cp()
	demo_map()
}
/**
   第二章 顺序编程-基础 demo end
*/
/**
   第三章 面向对象编程 demo start
*/
func chat_three()  {
	demo_integer()
}
/**
   第三章 面向对象编程 demo end
*/

func chat_five()  {

}

/**
* bool 值的demo
 */
func demo_boolean() {
	var b bool
	b = (1 != 0)
	fmt.Println("Result : ", b)

	var a float32
	a = 10 % -3
	fmt.Println("Result : ", a)
}

/**
* 比较 int 值的大小， int32 不能和 int64 的对比
 */
func demo_int_equal() {
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
复数的取值及赋值
*/
func demo_complex() {
	var complex_value complex64
	complex_value = 12 + 3i

	fmt.Println("Result : real number = ", real(complex_value), " image number = ", imag(complex_value))
}

func demo_string() {
	var str string
	str = "Hello world"
	ch := str[0]
	fmt.Println("The length of \" %s \" is %d \n", str, len(str))
	fmt.Println("THe first character of \" %s \"  is %c .\n", str, ch)
}

func demo_foreach() {
	str := "Hello, 世界."
	n := len(str)
	for i := 0; i < n; i++ {
		ch := str[i]
		fmt.Println(i, ch)
	}

	for i, ch := range str {
		fmt.Println(i, ch)
	}
}

func demo_slice() {
	// 定义一个数组
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//基于数组创建一个数组切片
	var mySlice []int = myArray[:5]

	fmt.Println("Elements of myArray : ")

	for index, v := range myArray {
		fmt.Print(index, "==>", v, " ")
	}

	fmt.Println("\nElements of mySlice: ")

	for index, v := range mySlice {
		fmt.Print(index, "==>", v, " ")
	}

	fmt.Println()
}

func demo_slice_two() {
	mySlice := make([]int, 5, 10)

	fmt.Println("len(mySlice) : ", len(mySlice))
	fmt.Println("cap(mySlice) :", cap(mySlice))
}

func demo_slice_two_cp() {
	fmt.Println("demo slice two cp.")
	oldSlice := []int{1, 2, 3, 4, 5}
	newSlice := oldSlice[:4]

	for index, value := range newSlice {
		fmt.Println(index, " ==>", value)
	}
}

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func demo_map() {
	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)
	personDB["12345"] = PersonInfo{"123456", "Tom", "Room  203 , ..."}
	personDB["1"] = PersonInfo{"1", "Jack", "Room 101,..."}

	person, ok := personDB["1234"]

	if ok {
		fmt.Println("Found person", person.Name, "with ID 1234.")
	} else {
		fmt.Println("Did not find person with ID 1234.")
	}
}

type Integer int

func (a Integer) Less(b Integer) bool{
	return a < b
}

func demo_integer()  {
	var a Integer = 1
	if a.Less(2){
		fmt.Println(a , "Less 2")
	}
}