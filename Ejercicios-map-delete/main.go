package main

import "fmt"

func main() {
	//Ejercicio 1 y 2
	m1 := make(map[string][]string)
	m1[`bond_james`] = []string{`shaken, not stirred`, `martinis`, `fast cars`}
	m1[`moneypenny_jenny`] = []string{`intelligence`, `literature`, `computer science`}
	m1[`no_dr`] = []string{`cats`, `ice cream`, `sunsets`}
	m1[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}
	fmt.Printf("%#v\n", m1)

	for k, v := range m1 {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Println(i, v2)
		}
	}
	fmt.Println("----------------recor deleted-------------------")
	delete(m1, "fleming_ian")
	fmt.Println("----------------recor deleted-------------------")
	for k, v := range m1 {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Println(i, v2)
		}
	}

}
