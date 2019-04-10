package common

import (
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
)

type Funclib string


/*获取当前文件执行的路径*/
func GetCurPath() string {
    file, _ := exec.LookPath(os.Args[0])
    //得到全路径，比如在windows下E:\\golang\\test\\a.exe
    path, _ := filepath.Abs(file)
    rst := filepath.Dir(path)
    return rst
}

/**
* 判断所给路径文件/文件夹是否存在
 */
func Exists(path string) bool {
    _,err := os.Stat(path)
    fmt.Println(err)
    if err != nil {
        if os.IsExist(err) {
            return true
        }
        return false
    }
    return true
}

/**
判断所给路基是否为文件夹
 */
func IsDir(path string) bool {
    s,err := os.Stat(path)
    if err != nil{
        return false
    }
    return  s.IsDir()
}

//判断所给路径是否为文件
func isFile(path string) bool {
    return  !IsDir(path)
}