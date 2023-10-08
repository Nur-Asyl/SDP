package main

import "fmt"

type Hero struct {
	attackSkills map[string]AttackStrategy
}

func NewHero() *Hero {
	fmt.Println("Pudge... Rooaar! M..ea..t...")
	return &Hero{
		attackSkills: make(map[string]AttackStrategy),
	}
}

func (h Hero) setAttack(attackName string, strategy AttackStrategy) {
	h.attackSkills["attackName"] = strategy
}

func (h Hero) makeAttack(attackStrategy string) {
	attackSkill := h.attackSkills[attackStrategy]

	attackSkill.Attack()
}
