package main

import "fmt"

func swap(x, y string, c int) (string, string, int) {
	return y, x, c
}

func sumamongus(valor int) (resultado1, resultado2 int) {
	valor = 5

	resultado1 = valor * 2

	resultado2 = valor + 2

	return
}

var y = 42
var z = "Amongus"
var x = `Amongs "Amongus"`

func main() {

	var c int
	var d int
	a, b, c := swap("hello", "world", 2)

	fmt.Println(a, b, c)
	fmt.Println(sumamongus(d))

	fmt.Println(y)

	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(x)
	fmt.Printf("%T\n", x)

}
