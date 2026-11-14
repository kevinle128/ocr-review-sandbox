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
