package main

import (
	"fmt"

	card "github.com/aQuaYi/Go-Patterns/DesignPatterns/factory-method/IDCard"
	"github.com/aQuaYi/Go-Patterns/DesignPatterns/factory-method/television"
)

func main() {
	factory := card.NewFactory()
	card := factory.Create("Alice")
	card.Use()

	fmt.Println("")
	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
	fmt.Println("")

	factory = television.NewFactory()
	tv := factory.Create("No.001")
	tv.Use()

}
