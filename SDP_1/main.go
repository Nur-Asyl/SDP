package main

import "fmt"

func main() {

	fmt.Println("CREATING HERO:")

	Pudge := NewHero()

	Pudge.setAttack("Hook", NewHookAttack())
	Pudge.setAttack("Root", NewRootAttack())
	Pudge.setAttack("Ultimate", NewUltimateDismemberAttack())

	fmt.Println("\nFIRST SKILL:")

	Pudge.makeAttack("Hook")

	fmt.Println("\nSECOND SKILL:")

	Pudge.makeAttack("Root")

	fmt.Println("\nULTIMATE SKILL:")

	Pudge.makeAttack("Ultimate")
}
