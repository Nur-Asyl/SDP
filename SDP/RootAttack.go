package main

import "fmt"

type RotAttack struct{}

func NewRotAttack() *RotAttack {
	return &RotAttack{}
}

func (r RotAttack) Attack() {
	fmt.Println("Its smell and melting...")
}
