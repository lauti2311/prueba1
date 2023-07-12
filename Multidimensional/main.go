package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Martini", "Chocolate"}
	jm := []string{"Jenny", "MoneyPenny", "Guiness", "Wolverine Tracks"}
	fmt.Println(jb)
	fmt.Println(jm)

	xp := [][]string{jb, jm}
	fmt.Println(xp)

}
