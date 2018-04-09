package autotesting

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		input          string
		expectedOutput string
	}{
		{"", ""},
		{"a", "a"},
		{"ab", "ba"},
		{"abc", "cba"},
		{"abcd", "dcba"},
		{"hé", "éh"},
		{"Hello, 世界", "界世 ,olleH"},
	}

	for _, c := range cases {
		output := Reverse(c.input)
		if output != c.expectedOutput {
			t.Errorf("incorrect output for %s: expected %s but got %s", c.input, c.expectedOutput, output)
		}
	}
}
