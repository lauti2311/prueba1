package main

import (
	"fmt"
)

// Short declaration operator
// Declaro una variable y le asigno un valor
// Utilizar := Para declarar variables dentro de funciones, tiene alcance desde la declaracion hasta el final de la función
// Una variable sin una inicializacion es igual a 0
// False para booleans, 0 para integers, 0.0 para floats, "" vacios para strings,
// y nil para pointer, funciones, interfaces, slcies, channels y mapas.
// La variable ab la declaramos fuera de una funcion sin necesidad de darle un valor (puede ser int, float, string)
// Declarar e iniciar "var ab=2", por ejemplo, es incicializar la variable
var ab int
var az int

func main() {

	x := 4

	b := "Amongus"

	var a int = 23 //utilizzar la palabra var para poder declarar variables fuera de funciones, tiene alcance en todo el package

	var z int

	z = a + x

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(x)
	fmt.Println(z)
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
