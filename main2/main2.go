package main

import "fmt"

func swap(x, y string, c int) (string, string, int) {
	return y, x, c
}

func main() {

	c := 1
	a, b, c := swap("hello", "world", 2)

	fmt.Println(a, b, c)
}
