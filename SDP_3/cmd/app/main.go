package main

import (
	"fmt"
	"module/internal/domain"
)

func main() {
	fmt.Println("CREATING HERO:")

	pudge := domain.NewWarrior()

	buff := &domain.Buff{}

	speedBuff := &domain.SpeedBuff{
		Buff: buff,
	}

	strengthAndSpeedBuff := &domain.StrengthBuff{
		Buff: speedBuff,
	}

	pudge.GetBuff(strengthAndSpeedBuff)

	pudge.CheckBuffs()

	pudge.SetAttack("Hook", domain.NewHookAttack())
	pudge.SetAttack("Root", domain.NewRootAttack())
	pudge.SetAttack("Ultimate", domain.NewUltimateDismemberAttack())
	m := &domain.MeteorCast{}
	mca := domain.MeteorCastAdapter{M: *m}
	pudge.SetAttack("Meteor", mca)
	b := &domain.BeemLaserCast{}
	blca := domain.BeemLaserCastAdapter{B: *b}
	pudge.SetAttack("Beem Laser", blca)

	fmt.Println("\nFIRST SKILL:")

	pudge.Attack("Hook")

	fmt.Println("\nSECOND SKILL:")

	pudge.Attack("Root")

	fmt.Println("\nULTIMATE SKILL:")

	pudge.Attack("Ultimate")

	fmt.Println("Meteor Spell:")

	pudge.Attack("Meteor")

	fmt.Println("Beem Laser Spell:")

	pudge.Attack("Beem Laser")
}
