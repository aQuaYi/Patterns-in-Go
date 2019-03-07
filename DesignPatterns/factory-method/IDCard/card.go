package card

import (
	"fmt"
	"time"

	frame "github.com/aQuaYi/Go-Patterns/DesignPatterns/factory-method/fmframe"
)

// H is
const H = 1

func ttt() {
	fmt.Println(frame.ID)
	return
}

// IDCard 是具体的产品类
type IDCard struct {
	owner string
}

func newIDCard(owner string) *IDCard {
	return &IDCard{
		owner: owner,
	}
}

// GetOwner returns IDCard's owner
func (c *IDCard) GetOwner() string {
	return c.owner
}

// Use implements Producter interface
func (c *IDCard) Use() {
	fmt.Printf("%s 使用了 IDCard :%s\n", c.owner, time.Now())
}
