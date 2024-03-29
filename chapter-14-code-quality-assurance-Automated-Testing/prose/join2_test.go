package prose

import "testing"

func TestTwoElements2(t *testing.T) {
	list := []string{"apple", "orange"}
	if JoinWithCommas(list) != "apple and orange" {
		t.Error("didn't match expected value")
	}
}

func TestThreeElements2(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	if JoinWithCommas(list) != "apple, orange, and pear" {
		t.Error("didn't match expected value")
	}

}
