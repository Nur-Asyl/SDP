package main

import "fmt"

type HookAttack struct{}

func NewHookAttack() *HookAttack {
	return &HookAttack{}
}

func (h HookAttack) Attack() {
	fmt.Println("Hook, BAAAAM!")
}
