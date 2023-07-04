package main

import "fmt"

//arrays, slice, map, struct
func main() {

	n := [...]string{"Almond Biscotti Café",
		"Banana Pudding ",
		"Balsamic Strawberry (GF)",
		"Bittersweet Chocolate (GF)",
		"Blueberry Pancake (GF)",
		"Bourbon Turtle (GF)",
		"Browned Butter Cookie Dough",
		"Chocolate Covered Black Cherry (GF)",
		"Chocolate Fudge Brownie",
		"Chocolate Peanut Butter (GF)",
		"Coffee (GF)",
		"Cookies & Cream",
		"Fresh Basil (GF)",
		"Garden Mint Chip (GF)",
		"Lavender Lemon Honey (GF)",
		"Lemon Bar",
		"Madagascar Vanilla (GF)",
		"Matcha (GF)",
		"Midnight Chocolate Sorbet (GF, V)",
		"Non-Dairy Chocolate Peanut Butter (GF, V)",
		"Non-Dairy Coconut Matcha (GF, V)",
		"Non-Dairy Sunbutter Cinnamon (GF, V)",
		"Orange Cream (GF) ",
		"Peanut Butter Cookie Dough",
		"Raspberry Sorbet (GF, V)",
		"Salty Caramel (GF)",
		"Slate Slate Different",
		"Strawberry Lemonade Sorbet (GF, V)",
		"Vanilla Caramel Blondie",
		"Vietnamese Cinnamon (GF)",
		"Wolverine Tracks (GF)"}
	fmt.Println(n)
	fmt.Println(len(n))
	fmt.Printf("%T\n", n)
	{
		//Declaro una variable [7] int
		var ni [7]int
		//Asigno un valor a la posicion 0

		ni[0] = 42
		fmt.Printf("%#v \t\t\t\t\t %T\n", ni, ni)

		//Array literal

		ni2 := [4]int{55, 56, 57, 58}
		fmt.Printf("%#v \t\t\t\t\t %T\n", ni2, ni2)

		//Array literal
		//Have compiler count elements
		ns := [...]string{"Chocolate", "Vainilla", "Frutilla"}

		fmt.Printf("%#v \t %T\n", ns, ns)

		//Usamos la funcion builtin len

		fmt.Println(len(ni))
		fmt.Println(len(ni2))
		fmt.Println(len(ns))

	}
}
