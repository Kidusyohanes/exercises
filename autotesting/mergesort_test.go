package autotesting

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	cases := []struct {
		input          []int
		expectedOutput []int
	}{
		{
			[]int{3, 1, 2},
			[]int{1, 2, 3},
		},
	}

	for _, c := range cases {
		output := MergeSort(c.input)
		if !reflect.DeepEqual(output, c.expectedOutput) {
			t.Errorf("incorrect output for %v: expected %v but got %v", c.input, c.expectedOutput, output)
		}
	}
}
