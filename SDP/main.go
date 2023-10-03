package main

import "fmt"

func main() {

	fmt.Println("CREATING HERO:")

	Pudge := NewHero()

	Pudge.setAttack("Hook", NewHookAttack())
	Pudge.setAttack("Root", NewRotAttack())

	fmt.Println("\nFIRST SKILL:")

	Pudge.makeAttack("Hook")

	fmt.Println("\nSECOND SKILL:")

	Pudge.makeAttack("Root")
}
