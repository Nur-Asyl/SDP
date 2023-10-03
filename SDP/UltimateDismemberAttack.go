package main

import "fmt"

type UltimateDismemberAttack struct {
}

func NewUltimateDismemberAttack() *UltimateDismemberAttack {
	return &UltimateDismemberAttack{}
}

func (u UltimateDismemberAttack) Attack() {
	fmt.Println("FRESH MEAT!!!")
}
