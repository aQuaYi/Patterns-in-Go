package card

import (
	"fmt"
	"time"
)

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
	fmt.Printf("%s 使用了一次 ID Card\n", c.owner)
	time.Sleep(time.Second)
}
