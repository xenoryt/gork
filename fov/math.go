package fov

type slope float64

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func calcSlope(newvar, newstatic int, oldvar, oldstatic int) slope {
	ds := float64(newstatic - oldstatic)
	dv := float64(newvar - oldvar)
	if ds != 0 {
		return slope(dv / ds)
	}
	return 0
}
