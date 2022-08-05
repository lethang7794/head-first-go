package prose3

import (
	"fmt"
	"testing"
)

type testData struct {
	list []string
	want string
}

func TestJoinWithCommas(t *testing.T) {
	tests := []testData{
		{
			list: []string{"my parents"},
			want: "my parents",
		},
		{
			list: []string{"apple", "orange"},
			want: "apple and orange",
		},
		{
			list: []string{"apple", "orange", "pear"},
			want: "apple, orange, and pear",
		},
	}
	for _, test := range tests {
		got := JoinWithCommas(test.list)
		if got != test.want {
			t.Errorf(errorString(test.list, got, test.want))
		}
	}
}

func errorString(list []string, got string, want string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, got, want)
}
