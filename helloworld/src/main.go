/**
* 注释风格与 C/C++ 保持一致，GO不要钱每个语句后边都加入分号,
* 左花括号 不能另起一行，否则会报编译错误 syntax error : unexpected semiconlon or newline before {
 */

//该代码所属的包名为 main;包是Go语言里的最基本的分发单位，也是工程管理中的依赖关系的体现
//要生产 GO可执行的程序，必须建立一个名为 main 的包，并且需要在包中包含一个叫main()的函数，是执行程序的执行起点
package main

//导入改程序需要的依赖包，不要在源代码中未使用的包，否则会报编译错误
import (
	"chat-eight"
	"fmt"
	"goenvsystem"
	"time"
)

// main()函数不能带任何参数，也不能定义返回值，
//命令行传入的参数是在 os.Args 变量中保存

func main() {
	/**
	  打印当前时间
	*/
	now()
	fmt.Println("Go,Lang,Hello,World!")
	_, _, last_name := get_name()
	fmt.Println(last_name)
	goenvsystem.GetSystemGoenv()
	//chat_five.ChatFive()
	chat_eight.ChatEight()
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
func now() {
	fmt.Println("当前时间 : ", time.Now())
}
