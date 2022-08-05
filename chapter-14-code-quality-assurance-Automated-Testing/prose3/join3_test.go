package prose3

import (
	"fmt"
	"testing"
)

func TestOneElements3(t *testing.T) {
	list := []string{"my parents"}
	want := "my parents"
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf(errorString(list, got, want))
	}
}

func TestTwoElements3(t *testing.T) {
	list := []string{"apple", "orange"}
	want := "apple and orange"
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf(errorString(list, got, want))
	}
}

func TestThreeElements3(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	want := "apple, orange, and pear"
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf(errorString(list, got, want))
	}
}

func errorString(list []string, got string, want string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, got, want)
}
