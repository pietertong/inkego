package chat_five

import	(
	"fmt"
	"os"
	"common"
)

//var typecommon common.Funclib

func main()  {
	s1 := pipe(func() int {
		return 100
	})
	s2 := format(func(s string, x, y int) string {
		return fmt.Sprintf(s, x, y)
	},"%d,%d",102,103)

	fmt.Println(fmt.Sprintf("sprint ln : %d,%d", 10,20));
	fmt.Println(s1,s2)

	fmt.Println(os.Getenv("GOPATH"))
}

func pipe(ff func() int) int {
	return ff()
}


type FormatFunc func(s string, x, y int) string

func format(ff FormatFunc, s string, x, y int) string {
	return ff(s, x, y)
}