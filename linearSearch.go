package main

func linearSearch(a []int, key int) int {
	for i := 0; i < len(a); i++ {
		if a[i] == key {
			return i
		}
	}
	return -1
}
