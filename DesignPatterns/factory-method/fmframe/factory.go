package frame

import "fmt"

// Factoryer 规定了工厂的接口
type Factoryer interface {
	Create(string) Producter
}

// Factory 是用于规范产品过程的工厂
type Factory struct {
	CreateProduct   func(string) Producter
	RegisterProduct func(Producter)
}

// Create 规范了产品的生产过程
func (f Factory) Create(name string) Producter {
	fmt.Println("----")
	fmt.Println("开始生产")
	fmt.Print("  第 1 步：")
	p := f.CreateProduct(name)
	fmt.Print("  第 2 步：")
	f.RegisterProduct(p)
	fmt.Println("完成生产")
	fmt.Println("----")
	return p
}
