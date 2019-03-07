package card

import (
	frame "github.com/aQuaYi/Go-Patterns/DesignPatterns/factory-method/fmframe"
)

// IDCardFactory 是具体的工厂类，用于定义生产的细节
type IDCardFactory struct {
	frame.Factory
	owners []string
}

// NewIDCardFactory 用于创建 IDCardFactory 对象
func NewIDCardFactory() *IDCardFactory {
	f := &IDCardFactory{}

	f.CreateProduct = func(owner string) frame.Producter {
		p := newIDCard(owner)
		return p
	}

	f.RegisterProduct = func(p frame.Producter) {
		c, _ := p.(*IDCard)
		f.owners = append(f.owners, c.GetOwner())
	}

	return f
}
