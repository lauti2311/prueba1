package main

import "fmt"

func main() {

	//apend, agregar valores a un slice
	xi := []int{41, 42, 43}

	fmt.Println(xi)
	fmt.Println("----------------")
	//variadic parameter

	xi = append(xi, 45, 46, 47, 48, 49, 999)

	fmt.Println(xi)
	fmt.Println("---------------")

	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)

	y := []int{234, 456, 678, 987}
	x = append(x, y...)
	fmt.Println(x)

}
