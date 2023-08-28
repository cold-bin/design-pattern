// @author cold bin
// @date 2023/8/28

package singleton

/*
	单例模式的饿汉式:项目编译运行的初期就已经初始化了对象
*/

type Logger interface {
	print(...any) error
}

type log struct{}

func (l *log) print(...any) error {
	return nil
}

var log_ *log

// 只要导入包，编译时便会初始化这个对象
func init() {
	log_ = &log{}
}

func GetInstance() Logger {
	return log_
}
