package domain

type MeteorCastAdapter struct {
	M MeteorCast
}

func (mca MeteorCastAdapter) attack() {
	mca.M.Cast()
}

type BeemLaserCastAdapter struct {
	B BeemLaserCast
}

func (blca BeemLaserCastAdapter) attack() {
	blca.B.Cast()
}
