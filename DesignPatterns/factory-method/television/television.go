package television

import (
	"fmt"
	"time"
)

// Television 是具体的产品类
type Television struct {
	sn string // serial number
}

func newTelevision(number string) *Television {
	return &Television{
		sn: number,
	}
}

// GetOwner returns Television's owner
func (t *Television) GetOwner() string {
	return t.sn
}

// Use implements Producter interface
func (t *Television) Use() {
	fmt.Printf("%s 播放了一次\n", t.sn)
	time.Sleep(time.Second)
}
