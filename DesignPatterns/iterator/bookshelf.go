package main

type bookShelf struct {
	books []book
}

func newBookShelf() *bookShelf {
	return &bookShelf{
		books: make([]book, 0, 256),
	}
}

func (bs *bookShelf) append(b book) {
	bs.books = append(bs.books, b)
}

func (bs *bookShelf) iterator() *bookShelfIterator {
	return &bookShelfIterator{
		bookShelf: bs,
	}
}
