package frame

import "fmt"

// Factoryer 规定了工厂接口
type Factoryer interface {
}

// Factory 是用于规范产品过程的工厂
type Factory struct {
	CreateProduct   func(string) Producter
	RegisterProduct func(Producter)
}

// Create 规范了产品的生产过程
func (f Factory) Create(name string) Producter {
	fmt.Printf("开始生产: %s", name)
	p := f.CreateProduct(name)
	f.RegisterProduct(p)
	fmt.Printf("完成生产: %s", name)
	return p
}
