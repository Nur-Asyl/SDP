package domain

import "fmt"

type MeteorCast struct{}

func NewMeteorCast() *MeteorCast {
	return &MeteorCast{}
}

func (m MeteorCast) Cast() {
	fmt.Println("^&*%^&*&^(&^&%$%&^*(_(*   Meteor!, BAAAAM!")
}
