package main

import (
	"fmt"
	"patterns/creational/factory/abstract"
	"patterns/creational/factory/factories"
)

func main() {
	ak47, _ := factories.GetGun("ak47")
	musket, _ := factories.GetGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g abstract.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}
