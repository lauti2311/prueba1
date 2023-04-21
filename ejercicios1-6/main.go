package main

import "fmt"

var a int = 42
var b string = "Amongus"
var c bool = true

type amongus int
type amogus int

var d amongus
var f amogus
var e int

func main() {
	//Ej1, usando  la declaracion de variable corta asignar distintos valores a x, y, z
	//Luego printearlos de forma individual y conjunta
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println(x, y, z)

	//Ej 2 usar var para declarar 3 variables a nivel package, no se le asignan valores
	//a como int, b como string y z como bool
	//en el main printear los valores
	//El compilador asigna valores a las variables como se llaman?
	//Le da un 0 a int, un string vacio a b y un false a c (zero value)
	fmt.Println(a, b, c)

	//EJ 3 asignar un valor a las variables var a=42 b="Amongus" c= true
	//Luego en el main Usar sprint para printear todo en un string, usar la variable s

	s := fmt.Sprintf("%v\t%v\t%v", a, b, c)
	fmt.Println(s)

	//EJ 4 Crear mi propio tipo de valor indentificado con int
	//Crear una variable con el owntype usando var
	//En el main printear el valor y tipo de la variable
	//asignarle 42 a la variable con =
	//Printear el valor de nuevo

	fmt.Println(d)
	fmt.Printf("%T\n", d)
	d = 42
	fmt.Println(d)

	//EJ 5 usar var para crear una variable e y debe ser el mismo tipo que la variable d
	//f = amogus int por lo tanto e= int
	//En el main printear el valor de f y el tipo y luego asignarle un valor
	//luego convertir el tipo del valor en f con el underlying type
	//Usar = para darle un valor a e
	//printear el valor de e y su tipo
	fmt.Println(f)
	fmt.Printf("%T\n", f)
	f = 22
	fmt.Println(f)
	e = int(f)
	fmt.Println(e)
	fmt.Printf("%T\n", e)
	//Ej 6 es el cuestionario
}
