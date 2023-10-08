package main

import "fmt"

type RootAttack struct{}

func NewRootAttack() *RootAttack {
	return &RootAttack{}
}

func (r RootAttack) Attack() {
	fmt.Println("Its smell and melting...")
}
