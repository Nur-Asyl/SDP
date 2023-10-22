package domain

type SpeedBuff struct {
	Buff IBuff
}

func (s *SpeedBuff) getDescription() string {
	buffDescription := s.Buff.getDescription()
	return buffDescription + "25% move speed\n"
}
