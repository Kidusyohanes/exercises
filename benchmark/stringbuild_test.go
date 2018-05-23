package benchmark

import (
	"strconv"
	"testing"
)

func TestConcat(t *testing.T) {
	cases := []struct {
		input    []string
		expected string
	}{
		{
			[]string{"a", "b", "c", "d"},
			"abcd",
		},
		{
			[]string{"The", "iSchool", "Is", "My", "School"},
			"TheiSchoolIsMySchool",
		},
		{
			[]string{},
			"",
		},
	}

	for _, c := range cases {
		actual := Concat(c.input...)
		if actual != c.expected {
			t.Errorf("incorrect output for Concat(%v): expected %s but got %s",
				c.input, c.expected, actual)
		}

		actual = Concat(c.input...)
		if actual != c.expected {
			t.Errorf("incorrect output for Concat2(%v): expected %s but got %s",
				c.input, c.expected, actual)
		}
	}
}

func genParts(nParts int) []string {
	parts := make([]string, 0, nParts)
	for i := 0; i < nParts; i++ {
		parts = append(parts, strconv.Itoa(i))
	}
	return parts
}
