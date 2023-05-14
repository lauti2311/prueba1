package main

import "fmt"

func main() {
	//For init; condition; post {}
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i := 0; i <= 10; i++ {
		fmt.Printf("The outer loop: %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t The inner loop: %d\n", j)

		}
	}

	x := 1
	for x < 10 {
		fmt.Println(x)
		x++

	}
	fmt.Println("Done")
	z := 1
	for {
		if z > 9 {
			break
		}
		fmt.Println(z)
		z++
	}
}
