package main

import "fmt"

func main() {
	a := Queue{}
	a.push(1)
	fmt.Println(a.pop())
	a.push(2)
	fmt.Println(a.pop())
	a.push(3)
	a.push(4)
	a.push(5)
	fmt.Println(a.pop())
	a.push(6)
	a.push(7)
	a.push(8)
	fmt.Println(a.pop())
	fmt.Println(a.pop())
	fmt.Println(a.pop())
	fmt.Println(a.pop())
	fmt.Println(a.pop())
}
