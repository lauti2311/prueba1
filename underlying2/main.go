package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []float64{3, 1, 4, 2}

	fmt.Println(mediaOne(n))
	fmt.Println("After mediaOne", n)

	y := []float64{3, 1, 4, 2}

	fmt.Println(mediaTwo(y))
	fmt.Println("After mediaTwo", y)

	//uses the same underlying array
	//everything is pass by value in go
	//the value is being passed into this func
	//and then assigned to x

}

func mediaOne(x []float64) float64 {
	sort.Float64s(x)
	i := len(x) / 2
	if len(x)%2 == 1 {
		return x[i/2]

	}
	return (x[i-1] + x[i]) / 2

}

func mediaTwo(x []float64) float64 {
	//allocate a new underlying array

	n := make([]float64, len(x))
	copy(n, x)

	sort.Float64s(n)
	i := len(n) / 2

	if len(n)%2 == 1 {
		return n[i/2]
		//when you divide
		//you get the whole number quotient
		//without the reminder nodulus
	}
	return (n[i-1] + n[i]) / 2

}
