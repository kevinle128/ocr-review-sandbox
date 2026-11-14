package sandbox

import "testing"

func TestClamp(t *testing.T) {
	tests := []struct {
		name      string
		v, lo, hi int
		want      int
	}{
		{"inside", 5, 0, 10, 5},
		{"below", -5, 0, 10, 0},
		{"above", 15, 0, 10, 10},
		{"at lo", 0, 0, 10, 0},
		{"at hi", 10, 0, 10, 10},
	}
	for _, tt := range tests {
		if got := Clamp(tt.v, tt.lo, tt.hi); got != tt.want {
			t.Errorf("%s: Clamp(%d, %d, %d) = %d, want %d", tt.name, tt.v, tt.lo, tt.hi, got, tt.want)
		}
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		xs   []int
		want int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"single", []int{7}, 7},
		{"multi", []int{7, 1, 2}, 10},
		{"negative", []int{-3, 1, 2}, 0},
	}
	for _, tt := range tests {
		if got := Sum(tt.xs); got != tt.want {
			t.Errorf("%s: Sum(%v) = %d, want %d", tt.name, tt.xs, got, tt.want)
		}
	}
}
