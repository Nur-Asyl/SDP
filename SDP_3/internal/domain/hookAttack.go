package domain

import "fmt"

type HookAttack struct{}

func NewHookAttack() *HookAttack {
	return &HookAttack{}
}

func (h HookAttack) attack() {
	fmt.Println("Hook, BAAAAM!")
}
