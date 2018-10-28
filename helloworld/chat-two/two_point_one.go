package chat_two

/**
2.1 变量
*/

/**
2.1.1 变量声明

Go 语言的变量声明方式 与 C 和 C++ 语言有明显不同
引入关键字 var ,而类型放在 变量名称之后
*/
var v1 int // 声明变量不需要 分号作为结束符
var v2 string
var v3 [10]int //数组
var v4 []int   // 切片
var v5 struct {
	f int
}
var v6 *int
var v7 map[string]int //map ,key为string类型，value 为 int 类型
var v8 func(a int) int

/**
var 关键字的另一种用法是可以将若干个需要什么的变量放置在一起，免得程序员需要重复写 var 关键字，如下
*/
var (
	v9  int
	v10 string
)

/**
2.1.2 变量初始化及赋值
*/

//var 关键字可以保留，但不再是必要的元素，如下所示

var v11 int = 10 //正确方式
var v12 = 10     //正确的方式，可以自动推导出 v2 的类型

//age := 10 //正确的方式，可以自动推导出 v13 的类型

/**
2.1.3 匿名变量,Python 也有类似的使用方式
*/
func GetName() (firstName, lastName, nickName string) {
	return "May", "Chan", "Chibi Maruko"
}

//_, _, nickName := GetName()
