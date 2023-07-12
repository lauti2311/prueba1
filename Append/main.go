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

}
