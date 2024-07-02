package main

import "fmt"

func multiplicationTable() {
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print(i * j)
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}
}
