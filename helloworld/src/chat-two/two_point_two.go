package chat_two

/**
2.2.1 常量定义
*/

const PI float64 = 3.141592666358979323846
const zero = 0.0 //无类型浮点常量
const (
	size int64 = 1024
	eof        = -1 //无类型整型常量
)

const u, v float32 = 0, 3    // u = 0,v = 3 ,常量的多重赋值
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

/**
2.2.3 预定义常量
*/

/**
Go 语言预定义了这些常量: true,false, iota;
*/

/**
iota 比较特殊，可以被认为是一个可被编译器修改的常量，
在每一个const 关键字出现时被重置为 0，
然后在下一个const 出现之前，每出现一次 iota,其所代表的数字就会自动增加 1
*/

//定义一组 常量
const ( //iota 被重设 为 0
	c0 = iota //c0 = 0
	c1 = iota //c1 = 1
	c2 = iota //c2 = 2
)

const (
	c3 = 1 << iota // c3 = 1 //(iota 在每个 const 开头被重设为 0)
	c4 = 1 << iota // c4 = 2
	c5 = 1 << iota // c5 = 4  位运算
)

const (
	u1         = iota * 42 // u == 0
	u2 float64 = iota * 42 // u2 = 42.0
	w          = iota * 42 // w = 84
)

const x = iota // x = 0 (因为 iota 又被重设为 0 了)
const y = iota // y = 0 (同上)

/**
2.2.4 枚举
*/

/**
枚举指 一系列的相关常量，
可以用在 const 后跟一对圆括号的方式定义一组常量，这种定义法在GO语言中常用于定义枚举值
GO 语言并不支持 enum 关键字
*/

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	numberOfDays // 这个常量没有导出
)

/**
同 GO 语言的其他符号 (symbol) 意义，以大写字母开头的常量在包外可见
上面例子中， numberOfDays 为包内私有，其他符号则可被其他包访问
*/
