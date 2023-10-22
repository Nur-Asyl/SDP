package domain

import "fmt"

type BeemLaserCast struct{}

func NewBeemLaserCast() *BeemLaserCast {
	return &BeemLaserCast{}
}

func (m BeemLaserCast) Cast() {
	fmt.Println("pshhhhiiiiiiiiiiiiiiiiiik, PSHIUYYY")
}
