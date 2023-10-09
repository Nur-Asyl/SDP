package domain

type StrengthBuff struct {
	Buff IBuff
}

func (s *StrengthBuff) getDescription() string {
	buffDescription := s.Buff.getDescription()
	return buffDescription + "10% hp, 10% damage\n"
}
