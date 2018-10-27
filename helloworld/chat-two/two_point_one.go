package chat_two

/**
Go 语言的变量声明方式 与 C 和 C++ 语言有明显不同

引入关键字 var ,而类型放在 变量名称之后
*/
var v1 int; // 声明变量不需要 分号作为结束符
var v2 string;
var v3 [10]int  //数组
var v4 []int  // 切片
var v5 struct{
    f int
}
var v6 *int
var v7 map[string] int //map ,key为string类型，value 为 int 类型
var v8 func(a int) int
/**
var 关键字的另一种用法是可以将若干个需要什么的变量放置在一起，免得程序员需要重复写 var 关键字，如下
*/
var (
    v9 int
    v10 string
)


/**
变量初始化
*/

//var 关键字可以保留，但不再是必要的元素，如下所示

var v11 int = 10 //正确方式
var v12 = 10 //正确的方式，可以自动推导出 v2 的类型

//age := 10 //正确的方式，可以自动推导出 v13 的类型


/**
常量定义
*/

const PI float64 = 3.141592666358979323846
const zero =0.0 //无类型浮点常量
const (
 size int64 = 1024
 eof = -1   //无类型整型常量
)

const u, v float32 = 0,3 // u = 0,v = 3 ,常量的多重赋值
const a, b, c = 3, 5, "foot" // a = 3, b = 5, c ="foot" ,无类型整型和字符串常量
/*
GO 的常量定义，可以限定常量类型，但不是必需的。
如果定义常量时没有指定类型，那么他与字母常量一样，是无类型的常量
*/
/**
常量定义的右值，也可以是一个在编译期运行的常量表达式
*/
const mask = 1 << 3

/**

常量定义的赋值是一个编译期行为，所以油脂不能出现任何需要运行期才能得出结果的表达式，比如
*/
//const Home = os.Getenv("home")
// 原因很简单，os.GetEnv()只有运行期才能知道返回结果，在编译期并不能确定，所以，无法作为常量定义的右值

