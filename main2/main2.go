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

func main() {

	var c int
	var d int
	a, b, c := swap("hello", "world", 2)

	fmt.Println(a, b, c)
	fmt.Println(sumamongus(d))

}
