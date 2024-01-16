package simpleinterest

func Calc(p, r, t float64) (interest float64) {
	interest = p * (r / 100) * t
	return
}
