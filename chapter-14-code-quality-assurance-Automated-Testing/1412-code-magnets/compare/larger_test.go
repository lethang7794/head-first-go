package compare

import (
	"fmt"
	"testing"
)

func TestFirstLarger(t *testing.T) {
	a, b := 2, 1
	want := 2
	got := Large(a, b)
	if got != want {
		t.Errorf(errorString(a, b, got, want))
	}
}

func TestSecondLarger(t *testing.T) {
	a, b := 4, 8
	want := 8
	got := Large(a, b)
	if got != want {
		t.Errorf(errorString(a, b, got, want))
	}
}

func errorString(a int, b int, got int, want int) string {
	return fmt.Sprintf("Larger(%d, %d) = %d, want %d", a, b, got, want)
}
