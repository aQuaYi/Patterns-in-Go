package main

type contexter interface {
	setClock(int)
	changeState(stater)
	callSecurityCenter(string)
	recordLog(string)
}
