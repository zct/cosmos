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

func MaxInt(a int, b ...int) int {
	max := a
	for _, v := range b {
		if v > max {
			max = v
		}
	}
	return max
}
