package main

import (
	"fmt"
	"module/internal/domain"
)

func main() {
	fmt.Println("CREATING HERO:")

	pudge := domain.NewHero()

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

	fmt.Println("\nFIRST SKILL:")

	pudge.MakeAttack("Hook")

	fmt.Println("\nSECOND SKILL:")

	pudge.MakeAttack("Root")

	fmt.Println("\nULTIMATE SKILL:")

	pudge.MakeAttack("Ultimate")
}
