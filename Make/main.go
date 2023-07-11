package main

import "fmt"

func main() {
	si := []string{"a", "b", "c"}

	fmt.Println(si)
	//creamos un slice sin nada deentro pero con capacidad de 0 a 10
	xi := make([]int, 0, 10)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))

	xi = append(xi, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	fmt.Println(xi)
	fmt.Println("-----------")
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	fmt.Println("-----------")
	xi = append(xi, 10, 11, 12, 13)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))

}
