package main

func numberOfEntry(a []int, q int) int {
	num := 0
	for i := 0; i < len(a); i++ {
		if a[i] == q {
			num++
		}
	}
	return num
}
