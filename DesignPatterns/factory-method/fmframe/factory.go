package frame

import "fmt"

// Factory 是用于规范产品过程的工厂
type Factory struct {
	createProduct   func(string) interface{}
	registerProduct func(interface{})
}

// Create 规范了产品的生产过程
func (f Factory) Create(name string) interface{} {
	fmt.Printf("开始生产: %s", name)
	p := f.createProduct(name)
	f.registerProduct(p)
	fmt.Printf("完成生产: %s", name)
	return p
}
