package main

import "fmt"

func containsElement() {
	var n int
	fmt.Scan(&n)
	a := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	var q string
	fmt.Scan(&q)

	for i := 0; i < n; i++ {
		if a[i] == q {
			fmt.Println("Contains Element")
			return
		}
	}
	fmt.Println("Not Contains Element")
}
