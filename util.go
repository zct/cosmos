package cosmos

func MinInt(a int, b ...int) int {
	min := a
	for _, v := range b {
		if v < min {
			min = v
		}
	}
	return min
}
