package sandbox

import "testing"

func TestAverage(t *testing.T) {
	if _, err := Average(nil); err == nil {
		t.Fatal("want error for empty input")
	}
	if got, _ := Average([]int{2, 4}); got != 3 {
		t.Fatalf("got %d", got)
	}
}

func TestMax(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{nil, 0},
		{[]int{}, 0},
		{[]int{-7}, -7},
		{[]int{2, 5, 3}, 5},
		{[]int{9, 1, 2}, 9},
		{[]int{-5, -2, -9}, -2},
		{[]int{-3, 0, -1}, 0},
	}
	for _, c := range cases {
		if got := Max(c.in); got != c.want {
			t.Errorf("Max(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}
