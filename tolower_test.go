package stringutil

import "testing"

func TestToLower(t *testing.T) {
	for _, c := range []struct {
		in, want string
	}{
		{"HELLO", "hello"},
		{"HellO", "hello"},
		{"", ""},
	} {
		got := ToLower(c.in)
		if got != c.want {
			t.Errorf("ToLower(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
