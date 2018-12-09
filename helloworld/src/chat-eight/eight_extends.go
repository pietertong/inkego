package chat_eight

import "fmt"

type Camera struct {

}
type Phone struct {

}

func (c *Camera) TakeAPicture() string {
    return "take a picture."
}

func (p *Phone) Call() string {
    return "make call."
}

type CameraPhone struct {
    Camera
    Phone
}

func eight_extends_main() {
    cp := new(CameraPhone)
    fmt.Println("我们的新款拍照手机有多种功能：")
    fmt.Println("打开了相机功能：",cp.TakeAPicture())
    fmt.Println("电话来了：",cp.Call())
}
