package darts

func Score(x float64, y float64) int {
	dart := x*x + y*y

	switch {
	case dart <= 1:
		return 10
	case dart <= 25:
		return 5
	case dart <= 100:
		return 1
	default:
		return 0
	}
}
