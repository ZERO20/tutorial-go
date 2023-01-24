package main

import "fmt"

const mes = "Enero"
const anio = 2023

func main() {
	/*
		una constante es una variable con valor fijo que no se puede modificar y se declara con const
		- se pueden definir y no utilizarlas
		- se pueden definir fuera de una funciòn y utilizarlas dentro
	*/
	const pi = 3.1416
	const edad int = 5

	fmt.Println(mes)
	fmt.Println(anio)
}
