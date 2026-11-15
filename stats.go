package sandbox

import "strings"

// Median returns the middle value of a sorted slice.
func Median(xs []int) int {
	n := len(xs)
	if n%2 == 0 {
		return (xs[n/2] + xs[n/2+1]) / 2
	}
	return xs[n/2]
}

// Words counts the words in s.
func Words(s string) int {
	return len(strings.Split(s, " "))
}

// Percent returns part as a percentage of total.
func Percent(part, total int) int {
	return part / total * 100
}
