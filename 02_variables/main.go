package main

import "fmt"

func main() {
	// declaraciòn sin valor
	var entero int    // default 0
	var cadena string // default ""
	var boleano bool  // default false

	// declaraciòn y asignaciòn
	var entero2 int = 5
	var cadena2 string = "Edgar"
	var boleano2 bool = true

	fmt.Println("Variables sin valor")
	fmt.Println(entero)
	fmt.Println(cadena)
	fmt.Println(boleano)

	fmt.Println("Variables con valor")
	fmt.Println(entero2)
	fmt.Println(cadena2)
	fmt.Println(boleano2)
}
