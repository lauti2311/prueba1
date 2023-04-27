package main

import (
	"fmt"
	"runtime"
)

var x bool
var d int
var e float64
var f int8 = -128 //int 8 almacena de -128 hasta 127, por ej -129 y 128 no funcionaria con int8

func main() {

	a := 7
	b := 2
	c := 42

	fmt.Println(x)
	x = true
	fmt.Println(x)

	fmt.Println(a == b) //Si pusiera a <= b (menor) mostraria verdadero a>=b (mayor) mostraria falso a != b (no es igual) true
	fmt.Print(c == b)
	//Numeric types
	d = 2
	e = 2.43212312

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)
	fmt.Printf("%T\n", f)

	fmt.Println(runtime.GOOS)   //Os donde corre
	fmt.Println(runtime.GOARCH) //Arquitectura donde corre
}
