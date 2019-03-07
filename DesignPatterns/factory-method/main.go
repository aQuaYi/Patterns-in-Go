package main

import (
	"fmt"

	card "github.com/aQuaYi/Go-Patterns/DesignPatterns/factory-method/IDCard"
	tv "github.com/aQuaYi/Go-Patterns/DesignPatterns/factory-method/television"
)

func main() {
	factory := card.NewIDCardFactory()

	cardA := factory.Create("A")
	cardB := factory.Create("B")
	cardC := factory.Create("C")

	cardA.Use()
	cardB.Use()
	cardC.Use()

	fmt.Println("")
	fmt.Println("-=-=-=-=-=-=-")
	fmt.Println("")

	tvFactory := tv.NewTelevisionFactory()

	tvA := tvFactory.Create("No.001")
	tvB := tvFactory.Create("No.002")
	tvC := tvFactory.Create("No.003")

	tvA.Use()
	tvB.Use()
	tvC.Use()
}
