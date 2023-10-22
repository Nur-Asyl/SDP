package domain

type WarriorAdapter struct {
	w Warrior
}

func (wa WarriorAdapter) Cast(magicItem IMagicItem) {
	wa.w.Attack(magicItem.Name())
}
