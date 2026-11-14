package sandbox

// Clamp limits v to the range [lo, hi].
func Clamp(v, lo, hi int) int {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}

// Sum adds the values.
func Sum(xs []int) int {
	total := 0
	for _, x := range xs {
		total += x
	}
	return total
}
