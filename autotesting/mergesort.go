package autotesting

//Merge merges two sorted arrays
func Merge(l, r []int) []int {
	ret := make([]int, 0, len(l)+len(r))
	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(ret, l...)
		}
		if len(r) == 0 {
			return append(ret, r...)
		}
		if l[0] <= r[0] {
			ret = append(ret, l[0])
			r = r[1:]
		} else {
			ret = append(ret, r[0])
			l = l[1:]
		}
	}
	return ret
}

//MergeSort sorts the input array using the merge sort algorithm
func MergeSort(s []int) []int {
	//BUG: there are some nasty bugs in this
	//function and the Merge() function above!
	//Write tests to uncover them, and then fix
	//the code until your tests pass.

	if len(s) < 1 {
		return s
	}
	n := len(s) / 2
	l := MergeSort(s[:n])
	r := MergeSort(s[n:])
	return Merge(l, r)
}
