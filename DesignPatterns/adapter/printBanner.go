package main

type printBanner struct {
	*banner
}

func newPrintBanner(b *banner) *printBanner {
	return &printBanner{
		banner: b,
	}
}

func (pb *printBanner) print() {
	pb.banner.show()
}

func (pb *printBanner) printWeak() {
	pb.banner.showWithParen()
}

func (pb *printBanner) printStrong() {
	pb.banner.showWithAster()
}
