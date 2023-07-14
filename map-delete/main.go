package main

import "fmt"

func main() {
	am := map[string]int{
		"Todd":     42,
		"Henry":    16,
		"Firulete": 23,
	}

	fmt.Println("La edad de Firulete era: ", am["Firulete"])
	fmt.Printf("%#v\n", am)

	an := make(map[string]int)
	an["Pepe"] = 28
	an["Jorge"] = 37
	an["Pedro"] = 78
	fmt.Println(an)
	fmt.Printf("%#v\n", an)
	fmt.Println(len(an))
	//BORRAR
	delete(an, "Jorge")

	fmt.Println("--- accessing keys that don't exist")
	delete(an, "Georage")      // won't panic
	fmt.Println(an["Georgey"]) // won't panic
	fmt.Println("------------------------")

	//For sobre un map}
	//k(key) son los "items", v los valores
	for k, v := range an {
		fmt.Println(k, v)
	}
	//_ es por si no quiero tener en cuenta la key se usa para indicar que se omite esa parte y el bucle for funcione
	for _, v := range an {
		fmt.Println(v)
	}
	//Solo para mostrar la key
	for k := range an {
		fmt.Println(k)
	}

	//For con un slice
	xi := []int{42, 43, 44}

	for i, v := range xi {
		fmt.Println(i, v)
	}
	for _, v := range xi {
		fmt.Println(v)
	}

	for i := range xi {
		fmt.Println(i)
	}
}
