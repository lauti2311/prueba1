package main

import "fmt"

func main() {
	xd := []string{"hello", "amongus"}

	fmt.Println(xd)

	xs := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)", "Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie", "Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)", "Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)", "Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough", "Raspberry Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different", "Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)"}

	fmt.Println(xs)
	fmt.Printf("%T\n", xs)

	for i, v := range xs {
		fmt.Printf(" %v -  %v\n", i, v)
	}

	{
		//for range slice

		xw := []string{"Almendra", "banana"}

		fmt.Println(xw)
		fmt.Printf("%T\n", xw)

		//blank indentifier (_)to not use a returned value

		for _, v := range xw {

			fmt.Printf("%v\n", v)

		}
		fmt.Println("----------------")
		fmt.Println(xw[0])
		fmt.Println(xw[1])
		fmt.Println(len(xw))

		for i := 0; i < len(xw); i++ {
			fmt.Println(xw[i])
		}
	}

}
