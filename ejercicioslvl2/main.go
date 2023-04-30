package main

import "fmt"

const (
	h int = 23

	i = "Hola"
)
const (
	m = 2019 + iota
	n = 2019 + iota
	o = 2019 + iota
	p = 2019 + iota
)

func main() {
	//Ej 1, printear un numero en decimal, binario y hex

	a := 23

	fmt.Printf("%d\t%b\t%#x", a, a, a)

	//Ej 2, usa los siguientes operadores, escribe expresiones y asigna sus valores a variables
	//==, <=, >=, !=, <, >
	b := (41 == 42)
	c := (21 <= 22)
	d := (34 >= 33)
	e := (31 != 32)
	f := (21 < 22)
	g := (10 > 18)

	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	//Ej 3 crear typed y untyped constants

	fmt.Println(h, i)

	//Ej 4. Asignar int a una variable, printearla en decimal, binario y hexadecimal
	//Mover ese valor de int a una posicion hacia la izquierda y asignarla a una variable
	//printear la variable en decimal binary and hex

	j := 8

	fmt.Printf("%d\t%b\t%#x\n", j, j, j)

	k := j << 1

	fmt.Printf("%d\t%b\t%#x\n", k, k, k)

	//Ej 5, crear una variable string usando un raw string literal, printearlo

	l := `A mi me gusta el amongus`

	fmt.Println(l)

	//Ej 6 usando iota, crear 4 constantes de los ultimos 4 años, printear los valores.
	fmt.Println(m, n, o, p)
}
