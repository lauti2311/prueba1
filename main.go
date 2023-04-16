package main

import "fmt"

func main() {

	fmt.Println("prueba")

	fmt.Println("Prueba")

	foo()

	fmt.Println("Algo mas")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}

	}

	bar()

}

func foo() {
	fmt.Println("Estoy en foo")
}

func bar() {
	fmt.Println("Salimos")
}

//control flow:
//1 Secuencial
//2  bucle iterativo
//condicional
