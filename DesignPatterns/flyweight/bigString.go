package main

import (
	"strings"
)

type bigString []*bigChar

func newBigString(str string) bigString {
	factory := newBigCharFactory()
	res := make([]*bigChar, len(str))
	for i := range res {
		res[i] = factory.getBigChar(str[i])
	}
	return res
}

func (bs bigString) String() string {
	var sb strings.Builder
	for i := range bs {
		sb.WriteString(bs[i].String())
	}
	return sb.String()
}
