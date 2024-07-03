package main

func intersectArrays(a []int, b []int) []int {
	dict1 := make(map[int]int)
	for _, v := range a {
		dict1[v]++
	}
	dict2 := make(map[int]int)
	for _, v := range b {
		dict2[v]++
	}
	var res []int
	for v, _ := range dict1 {
		if dict1[v] > 0 && dict2[v] > 0 {
			res = append(res, v)
		}
	}
	return res
}
