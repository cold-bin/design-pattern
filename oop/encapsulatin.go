// @author cold bin
// @date 2023/8/27

package oop

import "time"

/*
	封装：将实现细节隐藏，编写者有选择地对外暴露一些方法访问，
		 调用者不能拥有对内部结构直接修改的权限。
		 封装的核心就是访问权限控制
*/

type Wallet interface {
	GetBalance() uint64
	DecBalance(offset uint64)
	InrBalance(offset uint64)
}

type wallet struct {
	id      string
	created time.Time
	updated time.Time
	balance uint64
}

func (w *wallet) GetBalance() uint64 {
	return w.balance
}

// DecBalance 减少余额的同时也要更新时间，确保一致性
func (w *wallet) DecBalance(offset uint64) {
	w.balance -= offset
	w.updated = time.Now()
}

// InrBalance 增加余额的同时也要更新时间，确保一致性
func (w *wallet) InrBalance(offset uint64) {
	w.balance -= offset
	w.updated = time.Now()
}
