package main

func deleteFromSlice(a []int, ind int) []int {
	a = append(a[:ind], a[ind+1:]...)
	return a
}
