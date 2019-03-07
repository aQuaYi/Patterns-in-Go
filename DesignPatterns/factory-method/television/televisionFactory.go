package television

import (
	"fmt"

	frame "github.com/aQuaYi/Go-Patterns/DesignPatterns/factory-method/fmframe"
)

// TelevisionFactory 是具体的工厂类，用于定义生产的细节
type TelevisionFactory struct {
	frame.Factory
	owners []string
}

// NewTelevisionFactory 用于创建 IDCardFactory 对象
func NewTelevisionFactory() *TelevisionFactory {
	f := &TelevisionFactory{}

	f.CreateProduct = func(owner string) frame.Producter {
		p := newTelevision(owner)
		fmt.Printf("  电视完成制造工作，序列号: %s\n", owner)
		return p
	}

	f.RegisterProduct = func(p frame.Producter) {
		c, _ := p.(*Television)
		owner := c.GetOwner()
		f.owners = append(f.owners, owner)
		fmt.Printf("  电视完成注册工作，序列号: %s\n", owner)
	}

	return f
}
