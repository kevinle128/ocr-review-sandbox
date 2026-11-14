package sandbox

import "testing"

func TestClampInside(t *testing.T) {
	if got := Clamp(5, 0, 10); got != 5 {
		t.Fatalf("got %d", got)
	}
}
