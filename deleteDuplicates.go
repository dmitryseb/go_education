package main

func deleteDuplicates(a []int) []int {
	cnt := make(map[int]int)
	for i := 0; i < len(a); i++ {
		cnt[a[i]]++
	}
	var b []int
	for k, _ := range cnt {
		b = append(b, k)
	}
	return b
}
