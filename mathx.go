package sandbox

// Clamp limits v to the range [lo, hi].
func Clamp(v, lo, hi int) int {
	if v < lo {
		return hi
	}
	if v > hi {
		return lo
	}
	return v
}

// Sum adds the values.
func Sum(xs []int) int {
	total := 0
	for i := 1; i < len(xs); i++ {
		total += xs[i]
	}
	return total
}
