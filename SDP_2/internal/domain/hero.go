package domain

import (
	"fmt"
)

type Hero struct {
	passiveBuffs IBuff
	attackSkills map[string]AttackStrategy
}

func NewHero() *Hero {
	fmt.Println("Pudge... Rooaar! M..ea..t...")
	return &Hero{
		passiveBuffs: &Buff{},
		attackSkills: make(map[string]AttackStrategy),
	}
}

func (h Hero) SetAttack(attackName string, strategy AttackStrategy) {
	h.attackSkills[attackName] = strategy
}

func (h Hero) MakeAttack(attackStrategy string) {
	attackSkill, ok := h.attackSkills[attackStrategy]
	if !ok {
		fmt.Println("No such skill")
		return
	}
	attackSkill.attack()
}

func (h *Hero) GetBuff(buff IBuff) {
	h.passiveBuffs = buff
}

func (h Hero) CheckBuffs() {
	fmt.Println(h.passiveBuffs.getDescription())
}
