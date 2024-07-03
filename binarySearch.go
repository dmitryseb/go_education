package main

func binarySearch(a []int, q int) int {
	l := -1
	r := len(a)
	for r-l > 1 {
		m := (l + r) / 2
		if a[m] < q {
			l = m
		} else {
			r = m
		}
	}
	return r
}
