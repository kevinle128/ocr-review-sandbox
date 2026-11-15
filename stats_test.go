package sandbox

import "testing"

func TestMedian(t *testing.T) {
	tests := []struct {
		name string
		xs   []int
		want int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"single", []int{5}, 5},
		{"odd", []int{1, 2, 3}, 2},
		{"even 2", []int{2, 4}, 3},
		{"even 4", []int{1, 2, 3, 10}, 2},
		{"negative", []int{-5, -3, -1}, -3},
	}
	for _, tt := range tests {
		if got := Median(tt.xs); got != tt.want {
			t.Errorf("%s: Median(%v) = %d, want %d", tt.name, tt.xs, got, tt.want)
		}
	}
}

func TestWords(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"empty", "", 0},
		{"spaces only", "   ", 0},
		{"single", "hello", 1},
		{"multi", "a b c", 3},
		{"extra spaces", "  a   b ", 2},
		{"tabs and newlines", "a\tb\nc", 3},
	}
	for _, tt := range tests {
		if got := Words(tt.s); got != tt.want {
			t.Errorf("%s: Words(%q) = %d, want %d", tt.name, tt.s, got, tt.want)
		}
	}
}

func TestPercent(t *testing.T) {
	tests := []struct {
		name        string
		part, total int
		want        int
	}{
		{"zero total", 5, 0, 0},
		{"zero part", 0, 10, 0},
		{"half", 1, 2, 50},
		{"three quarters", 3, 4, 75},
		{"whole", 4, 4, 100},
		{"over", 6, 4, 150},
	}
	for _, tt := range tests {
		if got := Percent(tt.part, tt.total); got != tt.want {
			t.Errorf("%s: Percent(%d, %d) = %d, want %d", tt.name, tt.part, tt.total, got, tt.want)
		}
	}
}

func TestMean(t *testing.T) {
	tests := []struct {
		name string
		xs   []int
		want int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"single", []int{7}, 7},
		{"multi", []int{2, 4, 6}, 4},
		{"truncates", []int{1, 2}, 1},
		{"negative", []int{-3, -6}, -4},
		{"mixed", []int{-2, 2, 6}, 2},
	}
	for _, tt := range tests {
		if got := Mean(tt.xs); got != tt.want {
			t.Errorf("%s: Mean(%v) = %d, want %d", tt.name, tt.xs, got, tt.want)
		}
	}
}
