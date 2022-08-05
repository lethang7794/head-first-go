package arithmetic

import "testing"

func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Error("1 + 2 did not equal 3")
	}
}

func TestSubtract(t *testing.T) {
	if Subtract(3, 2) != 1 {
		t.Error("3 - 2 did not equal 1")
	}
}
