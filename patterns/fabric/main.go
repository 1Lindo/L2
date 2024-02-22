package main

import (
	"fmt"
	"main/pkg/fabricPkg"
)

func main() {
	ak47, _ := fabricPkg.GetGun("ak47")
	musket, _ := fabricPkg.GetGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g fabricPkg.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}
