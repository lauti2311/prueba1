package main

import "fmt"

func main() {
	//Usar composite literal:
	//crear un array que tenga 5 valoresd de tipo int
	//ingresar valores para cada posicion del index
	fmt.Println("EJERCICIO 1:")
	x := [5]int{}
	for i := 0; i < 5; i++ {
		x[i] = i

	}
	for i, v := range x {
		fmt.Printf("%v\t- %T\t - %v\n", v, v, i)
	}

	//ejercicio 2:
	/*
		Usar composite literal:
		Crear un slice de tipo innt
		asignar 10 valores: 42,43,44,45,46,47,48,49,50,51
		printear los valores con el valor y tipo
	*/
	fmt.Println("EJERCICIO 2 y 3:")
	xs := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i, v := range xs {
		fmt.Printf("%v \t - %T \t- %v\n", v, v, i)
	}

	fmt.Printf("%T \t %#v\n", x, x)
	fmt.Printf("%T \t %v\n", x, x)

	/*Ejercicio 3:
	mostrar con el slice anterior el siguiente orden
	[42 43 44 45 46]
	[47 48 49 50 51]
	[44 45 46 47 48]
	[43 44 45 46 47]
	*/

	fmt.Println(xs[:5])
	fmt.Println(xs[5:])
	fmt.Println(xs[2:7])
	fmt.Println(xs[1:6])
	/*
		Ejercicio 4:
		Follow these steps:
		start with this slice
		x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
		append to that slice this value
		52
		print out the slice
		in ONE STATEMENT append to that slice these values
		53
		54
		55
		print out the slice
		append to the slice this slice
		y := []int{56, 57, 58, 59, 60}
		print out the slice

	*/

	fmt.Println("EJERCICIO 4:")
	xa := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	xa = append(xa, 52)

	fmt.Println(xa)

	xa = append(xa, 53, 54, 55)

	fmt.Println(xa)

	y := []int{56, 57, 58, 59, 60}

	xa = append(xa, y...)

	fmt.Println(xa)

	/* Ejercicio 5
	To DELETE from a slice, we use APPEND along with SLICING. For this hands-on exercise, follow these steps:
	start with this slice
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	use APPEND & SLICING to get these values here which you should ASSIGN to the variable “x” and then print:
	[42, 43, 44, 48, 49, 50, 51]


	*/

	fmt.Println("EJERCICIO 5:")

	b := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(b)

	b = append(b[:3], b[6:]...)

	fmt.Println(b)
	/*
		For this exercise, do the following:
		Create a slice to store the names of all of the states in the United States of America.
		Use make and append to do this.
		Goal: do not have the array that underlies the slice created more than once.
		Print out
		the len
		the cap
		the values, along with their index position, without using the range clause.
		Here is a list of the 50 states:

		` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,


	*/
	fmt.Println("EJERCICIO 6:")

	c := make([]int, 50)

	fmt.Println(c)
	fmt.Println(len(c))
	fmt.Println(cap(c))

	c2 := make([]int, 0, 50)

	fmt.Println(c2)
	fmt.Println(len(c2))
	fmt.Println(cap(c2))

	fmt.Println("-----------------------")
	c[0] = 98
	c2 = append(c2, 99)

	fmt.Println(c)
	fmt.Println(len(c))
	fmt.Println(cap(c))

	fmt.Println(c2)
	fmt.Println(len(c2))
	fmt.Println(cap(c2))

	fmt.Println("-----------------------")

	ac := make([]string, 0, 50)

	fmt.Println(len(ac))
	fmt.Println(cap(ac))

	ac = append(ac, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)

	fmt.Println(len(ac))
	fmt.Println(cap(ac))

	for i := 0; i < len(ac); i++ {

		fmt.Println(ac[i], i)

	}

	fmt.Println("EJERCICIO 7:")
	/*
		Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:

		"James", "Bond", "Shaken, not stirred"
		"Miss", "Moneypenny", "I'm 008."

		Range over the records, then range over the data in each record.

	*/

	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mm := []string{"Miss", "Moneypenny", "I'm 008."}

	jm := [][]string{jb, mm}

	for i, v := range jm {
		fmt.Println(i, v)
		for a, b := range v {
			fmt.Println(a, b)
		}
	}

}
