package string

import "testing"

func Test(t *testing.T) {
	tests := []struct{ result, expected string }{
		{"Hello!", "!olleH"},
		{"Hello Dima!", "!amiD olleH"},
	}

	for _, v := range tests {
		got := Reverse(v.result)

		if got != v.expected {
			t.Errorf("Reverse(%q) == %q, expect %q", v.result, got, v.expected)
		}
	}
}
