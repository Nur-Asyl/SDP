package domain

import "fmt"

type UltimateDismemberAttack struct{}

func NewUltimateDismemberAttack() *UltimateDismemberAttack {
	return &UltimateDismemberAttack{}
}

func (u UltimateDismemberAttack) attack() {
	fmt.Println("FRESH MEAT!!!")
}
