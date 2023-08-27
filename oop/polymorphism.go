// @author cold bin
// @date 2023/8/27

package oop

/*
	多态是指代码可以根据类型的具体实现采取不同行为的能力。
	因为任何用户定义的类型都可以实现任何接口，所以通过不同实体类型对接口值方法的调用就是多态。
*/

import "fmt"

type Usbs interface {
	Start()
	Stop()
}

type Phones struct {
}

// Phone 实现 Usb 接口的方法
func (p Phones) Start() {
	fmt.Println("手机开始工作")
}

func (p Phones) Stop() {
	fmt.Println("手机停止工作")
}

type Cameras struct {
}

// 让Camera实现 Usb 的方法
func (cameras Cameras) Start() {
	fmt.Println("相机开始工作")
}

func (cameras Cameras) Stop() {
	fmt.Println("相机停止工作")
}

type Computer struct {
}

func (c Computer) Working(usbs Usbs) { // usb 变量会根据传入的实参，来判断到底是 Phone,还是 Camera
	//通过 usb 接口变量来调用 Start 和 Stop 方法
	usbs.Start()
	usbs.Stop()
}

func main_() {
	computer := Computer{}
	phones := Phones{}
	cameras := Cameras{}
	computer.Working(phones)
	computer.Working(cameras)
}
