package main

func merge(a []int, b []int) []int {
	result := make([]int, len(a)+len(b))
	j := 0
	for i := 0; i < len(a); i++ {
		for j < len(b) && b[j] < a[i] {
			result[i+j] = b[j]
			j++
		}
		result[i+j] = a[i]
	}
	for j < len(b) {
		result[len(a)+j] = b[j]
		j++
	}
	return result
}
