package main

import "fmt"

func reverseCounter() {
	var n int
	fmt.Scan(&n)
	for i := n; i > 0; i-- {
		fmt.Println(i)
	}
}
