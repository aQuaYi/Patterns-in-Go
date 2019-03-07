package card

import (
	"fmt"

	frame "github.com/aQuaYi/Go-Patterns/DesignPatterns/factory-method/fmframe"
)

// factory 是具体的工厂类，用于定义生产的细节
type factory struct {
	frame.Factory
	owners []string
}

// NewFactory 返回专门用来创建 *IDCard 的工厂实例
func NewFactory() frame.Factoryer {
	f := &factory{}

	f.CreateProduct = func(owner string) frame.Producter {
		p := newIDCard(owner)
		fmt.Printf("创建 %s 的 ID Card\n", owner)
		return p
	}

	f.RegisterProduct = func(p frame.Producter) {
		c, _ := p.(*IDCard)
		owner := c.GetOwner()
		f.owners = append(f.owners, owner)
		fmt.Printf("注册 %s 的 ID Card \n", owner)
	}

	return f
}
