package main

import (
	"fmt"
	"strings"
)

func main() {
	bs := newBigString("110")
	fmt.Println(bs)

	// 轻量的秘密在于共享数据
	// 所以，修改 bs[0] 后，bs[1] 中的 "1" 也被修改了。
	one := bs[0]
	one.fontData = strings.Replace(one.fontData, ".", "`", -1)
	fmt.Println(bs)

}
