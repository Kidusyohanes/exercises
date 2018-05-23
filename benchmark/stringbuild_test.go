package benchmark

import (
	"fmt"
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

// func BenchmarkConcat(b *testing.B) {
// 	parts := genParts(1000)
// 	for i := 0; i < b.N; i++ {
// 		Concat(parts...)
// 	}
// }
// func BenchmarkConcat2(b *testing.B) {
// 	parts := genParts(1000)
// 	for i := 0; i < b.N; i++ {
// 		Concat2(parts...)
// 	}
// }

func BenchmarkConcats(b *testing.B) {
	for nParts := 10; nParts < 100000; nParts *= 10 {
		parts := genParts(nParts)
		b.Run(fmt.Sprintf("Concat-%d", nParts), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Concat(parts...)
			}
		})
		b.Run(fmt.Sprintf("Concat2-%d", nParts), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Concat2(parts...)
			}
		})
	}
}
