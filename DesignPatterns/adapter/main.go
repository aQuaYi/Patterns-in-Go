package main

func main() {
	pb := newPrintBanner(newBanner("content"))

	// NOTICE:
	// print() 只接受 printer 接口
	// banner 只实现了 shower 接口
	// 通过创建新类 printBanner 作为适配器
	// 在 printBanner 实现 printer 接口时
	// 把具体的工作委托给 banner 的方法完成
	print(pb)
}
