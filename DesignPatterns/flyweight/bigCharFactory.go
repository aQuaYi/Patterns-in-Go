package main

import "sync"

var (
	once      sync.Once
	singleton *bigCharFactory
)

type bigCharFactory struct {
	pool map[byte]*bigChar
}

func newBigCharFactory() *bigCharFactory {
	once.Do(func() {
		if singleton == nil {
			singleton = &bigCharFactory{
				pool: make(map[byte]*bigChar, 10),
			}
		}
	})
	return singleton
}

func (f *bigCharFactory) getBigChar(char byte) *bigChar {
	bc, ok := f.pool[char]
	if !ok {
		bc = newBigChar(char)
		f.pool[char] = bc
	}
	return bc
}
