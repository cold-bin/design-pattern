// @author cold bin
// @date 2023/8/27

package oop

import "fmt"

/*
	继承可以提高代码复用率。可以将一些基类的基础函数或方法通过继承来获取，不必再重写，起到代码复用的目睹
*/

type product struct {
	name      string
	introduce string
	price     uint64
}

func (p *product) GetPrice() uint64 {
	return p.price
}

func (p *product) UpdateName(newName string) {
	p.name = newName
}

type phone struct {
	product
	power     uint8
	osVersion string
}

func (p phone) Show() {
	p.GetPrice()
	fmt.Println(p)
}
