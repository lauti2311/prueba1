package main

import "fmt"

var a int

type amongus int //Esto quiere decir que amongus va a ser un tipo int subyacente
var b amongus

func main() {
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	//conversion

	a = int(b) //Aunque b sea tipo amongus, al poner int delante lo transforma en un int
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
