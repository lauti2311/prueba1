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

	//Division, resto
	c := 41 % 40
	d := 41 / 40
	fmt.Println(c, d)
	a := 1
	for {
		a++
		if a > 100 {
			break
		}
		if a%2 != 0 {
			continue
		}
		fmt.Println(a)

	}

	fmt.Println("FIN")
	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t%#x\t%#U\n", i, i, i) //asci %valor, ascii %#U, %#x hexadecimal
	}

}
