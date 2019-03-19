package main

func main() {

	b1 := newStringDisplay("Hello, World.")
	b1.show()

	b2 := newSideBorder(b1, '#')
	b2.show()

	b3 := newFullBorder(b2)
	b3.show()

	b4 :=
		newSideBorder(
			newFullBorder(
				newSideBorder(
					newSideBorder(
						newFullBorder(
							newFullBorder(
								newSideBorder(
									newFullBorder(
										newStringDisplay("Good Luck!"),
									), '*',
								),
							),
						), '/',
					), '$'),
			), '#')
	b4.show()

	// #+----------------------+#
	// #|$/+----------------+/$|#
	// #|$/|+--------------+|/$|#
	// #|$/||*+----------+*||/$|#
	// #|$/||*|Good Luck!|*||/$|#
	// #|$/||*+----------+*||/$|#
	// #|$/|+--------------+|/$|#
	// #|$/+----------------+/$|#
	// #+----------------------+#

}
