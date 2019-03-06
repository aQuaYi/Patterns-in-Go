package main

type bookShelfIterator struct {
	*bookShelf
	nextIndex int
}

func (bsi *bookShelfIterator) hasNext() bool {
	return bsi.nextIndex < len(bsi.bookShelf.books)
}

func (bsi *bookShelfIterator) next() interface{} {
	book := bsi.bookShelf.books[bsi.nextIndex]
	bsi.nextIndex++
	return book
}
