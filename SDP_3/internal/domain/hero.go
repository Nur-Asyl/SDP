package domain

import (
	"fmt"
)

type Mage struct {
	passiveBuffs IBuff
	attackSkills map[string]AttackStrategy
}

func NewMage() *Mage {
	fmt.Println("I am Mage!")
	return &Mage{
		passiveBuffs: &Buff{},
		attackSkills: make(map[string]AttackStrategy),
	}
}

func (m Mage) SetAttack(attackName string, strategy AttackStrategy) {
	m.attackSkills[attackName] = strategy
}

func (m Mage) Cast(attackStrategy string) {
	attackSkill, ok := m.attackSkills[attackStrategy]
	if !ok {
		fmt.Println("No such skill")
		return
	}
	attackSkill.attack()
}

func (m *Mage) GetBuff(buff IBuff) {
	m.passiveBuffs = buff
}

func (m Mage) CheckBuffs() {
	fmt.Println(m.passiveBuffs.getDescription())
}

type Warrior struct {
	passiveBuffs IBuff
	attackSkills map[string]AttackStrategy
}

func NewWarrior() *Warrior {
	fmt.Println("I am Mage!")
	return &Warrior{
		passiveBuffs: &Buff{},
		attackSkills: make(map[string]AttackStrategy),
	}
}

func (w Warrior) SetAttack(attackName string, strategy AttackStrategy) {
	w.attackSkills[attackName] = strategy
}

func (w Warrior) Attack(attackStrategy string) {
	attackSkill, ok := w.attackSkills[attackStrategy]
	if !ok {
		fmt.Println("No such skill")
		return
	}
	attackSkill.attack()
}

func (w *Warrior) GetBuff(buff IBuff) {
	w.passiveBuffs = buff
}

func (w Warrior) CheckBuffs() {
	fmt.Println(w.passiveBuffs.getDescription())
}
