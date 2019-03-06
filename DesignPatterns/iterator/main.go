package main

func main() {
	bs := newBookShelf()
	bs.append(newBook("Around the World in 80 Days"))
	bs.append(newBook("Bible"))
	bs.append(newBook("Cinderella"))
	bs.append(newBook("Daddy-long-legs"))

	bsi := bs.iterator()

	// NOTICE: 通过 iterator 接口，printAll 与 bsi 实现了解耦
	printAll(bsi)
}
