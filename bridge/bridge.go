// @author cold bin
// @date 2023/9/2

package bridge

import "fmt"

// AlertMethod 告警方式的接口
type AlertMethod interface {
	SendAlert(message string)
}

// 具体的告警方式

type EmailAlert struct{}

func (e *EmailAlert) SendAlert(message string) {
	fmt.Println("通过邮件发送告警：", message)
}

type SMSAlert struct{}

func (s *SMSAlert) SendAlert(message string) {
	fmt.Println("通过短信发送告警：", message)
}

// AlertLevel 告警级别的接口
type AlertLevel interface {
	SetAlertMethod(method AlertMethod)
	Alert(message string)
}

// 具体的告警级别

type WarningAlert struct {
	method AlertMethod
}

func (w *WarningAlert) SetAlertMethod(method AlertMethod) {
	w.method = method
}

func (w *WarningAlert) Alert(message string) {
	w.method.SendAlert("[Warning] " + message)
}

type ErrorAlert struct {
	method AlertMethod
}

func (e *ErrorAlert) SetAlertMethod(method AlertMethod) {
	e.method = method
}

func (e *ErrorAlert) Alert(message string) {
	e.method.SendAlert("[Error] " + message)
}
