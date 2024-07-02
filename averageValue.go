package main

func averageValue(a []int) float64 {
	sum := 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return float64(sum) / float64(len(a))
}
