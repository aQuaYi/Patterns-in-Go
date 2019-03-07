package television

import (
	"fmt"
	"time"
)

// Television 是具体的产品类
type Television struct {
	owner string
}

func newTelevision(owner string) *Television {
	return &Television{
		owner: owner,
	}
}

// GetOwner returns Television's owner
func (t *Television) GetOwner() string {
	return t.owner
}

// Use implements Producter interface
func (t *Television) Use() {
	time.Sleep(time.Second)
	fmt.Printf("%s 播放了一秒钟 :%s\n", t.owner, time.Now())
}
