package main

import "fmt"

const a = 42
const b = 42.78
const c = "Amongus"
const (
	d = iota
	e = iota
	f = iota //Los pone como tipo int
	//Tambien puedo ponerle iota al primer valor, y los que le siguen, van a incrementar su valor
)

//Mediante iota vamos multiplicando el numero.
const (
	_ = iota
	//kb = 1024
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

//Si lo defino como int, float64 o string lo hago una constante de tipo
//Tambien puedo poner todo junto como un const () Dentro de los parentesis valores y demas
func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	x := 2
	fmt.Printf("%d\t\t%b\n", x, x)

	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)

	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", mb, mb)

}
