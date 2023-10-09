package domain

type IBuff interface {
	getDescription() string
}

type Buff struct{}

func (b *Buff) getDescription() string {
	return "BUFFS:\n"
}
