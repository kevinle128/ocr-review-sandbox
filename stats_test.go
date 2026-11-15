package sandbox

import "testing"

func TestMedianOdd(t *testing.T) {
	if got := Median([]int{1, 2, 3}); got != 2 {
		t.Fatalf("got %d", got)
	}
}
