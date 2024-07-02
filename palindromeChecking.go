package main

import "fmt"

func palindromeChecking() {
	var s string
	fmt.Scan(&s)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			fmt.Println(s + " is not a palindrome.")
			return
		}
	}
	fmt.Println(s + " is a palindrome.")
}
