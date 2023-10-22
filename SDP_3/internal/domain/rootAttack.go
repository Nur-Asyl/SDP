package domain

import "fmt"

type RootAttack struct{}

func NewRootAttack() *RootAttack {
	return &RootAttack{}
}

func (r RootAttack) attack() {
	fmt.Println("Its smell and melting...")
}
