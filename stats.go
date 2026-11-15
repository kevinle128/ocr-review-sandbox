package sandbox

import "strings"

// Median returns the middle value of a sorted slice, or 0 for an empty slice.
func Median(xs []int) int {
	n := len(xs)
	if n == 0 {
		return 0
	}
	if n%2 == 0 {
		return (xs[n/2-1] + xs[n/2]) / 2
	}
	return xs[n/2]
}

// Words counts the whitespace-separated words in s.
func Words(s string) int {
	return len(strings.Fields(s))
}

// Percent returns part as a percentage of total, or 0 when total is 0.
func Percent(part, total int) int {
	if total == 0 {
		return 0
	}
	return part * 100 / total
}
