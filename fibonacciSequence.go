package main

func fibonacciSequence(n int) []int {
	fibonacci := make([]int, n)
	if n > 0 {
		fibonacci[0] = 1
	}
	if n > 1 {
		fibonacci[1] = 1
	}
	for i := 2; i < n; i++ {
		fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
	}

	return fibonacci
}
