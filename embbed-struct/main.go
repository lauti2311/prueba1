package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}
type secretAgent struct {
	person
	itk   bool
	first string
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "bond",
			age:   32,
		},
		first: "ETHAN HAWK",
		itk:   true,
	}

	p2 := person{
		first: "Jenny",
		last:  "MoneyPenny",
		age:   27,
	}
	fmt.Println(sa1)
	fmt.Println(p2)
	fmt.Printf("%T \t %#v\n", sa1, sa1)

	fmt.Println(sa1.first, sa1.last, sa1.age)
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.itk, sa1.person)
	fmt.Println(sa1.first, sa1.person.first)
}
