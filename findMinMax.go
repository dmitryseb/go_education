package main

func findMinMax(a []int) (minVal int, maxVal int) {
	minVal = a[0]
	maxVal = a[0]
	for i := 1; i < len(a); i++ {
		minVal = min(minVal, a[i])
		maxVal = max(maxVal, a[i])
	}
	return
}
