package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	b := a

	c := make([]int, 6)

	copy(c, a)

	fmt.Println("a ", a)
	fmt.Println("b ", b)
	fmt.Println("c ", c)
	fmt.Println("------------")

	a[0] = 7

	fmt.Println("a ", a)
	fmt.Println("b ", b)
	fmt.Println("c ", c)
	fmt.Println("------------")
}
