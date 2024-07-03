package main

import "strings"

func checkAnagram(a string, b string) bool {
	cnt1 := make(map[uint8]int)
	a = strings.ToLower(a)
	for i := 0; i < len(a); i++ {
		cnt1[a[i]]++
	}
	b = strings.ToLower(b)
	cnt2 := make(map[uint8]int)
	for i := 0; i < len(b); i++ {
		cnt2[b[i]]++
	}
	for v, u := range cnt1 {
		if cnt2[v] != u {
			return false
		}
	}
	for v, u := range cnt2 {
		if cnt1[v] != u {
			return false
		}
	}
	return true
}
