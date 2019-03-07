package frame

import "fmt"

// FIXME: 修复此处内容
// // Factoryer 规定了工厂的接口
// type Factoryer interface {
// 	Creat(string) Producter
// }

// Factory 是用于规范产品过程的工厂
type Factory struct {
	CreateProduct   func(string) Producter
	RegisterProduct func(Producter)
}

// Create 规范了产品的生产过程
func (f Factory) Create(name string) Producter {
	fmt.Println("开始生产")
	p := f.CreateProduct(name)
	f.RegisterProduct(p)
	fmt.Println("完成生产")
	fmt.Println("----")
	return p
}
