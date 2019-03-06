package main

func main() {
	bs := newBookShelf()
	bs.append(newBook("Around the World in 80 Days"))
	bs.append(newBook("Bible"))
	bs.append(newBook("Cinderella"))
	bs.append(newBook("Daddy-long-legs"))

	bsi := bs.iterator()

	printAll(bsi)
}
