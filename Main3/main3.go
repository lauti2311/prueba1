package main

import "fmt"

var y string
var x = 10

func main() {
	fmt.Println("Zero value")

	//Declarar una variable de cierto tipo
	// y luego asignar un valor del tipo de la variable

	fmt.Println("Printing el valor de y", y, "fin")

	fmt.Printf("%T\n", y)

	y = "Amongus, Amogus Amongus"

	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(x)
	fmt.Printf("%T", x)  //Tipo
	fmt.Printf("%b", x)  //binario
	fmt.Printf("%x", x)  //hexadecimal
	fmt.Printf("%#x", x) //0 delante y hexadecimal

	//EL SPRINT,SPRINTF,SPRINTLN, es para asignar a una variable un string de lo que se ponga como parametro

}
