package main

import "fmt"

/*
Operador de direcciòn => & direcciòn de memoria
Operador de desreferenciaciòn => * valor de la direcciòn de memoria
*/
func main() {
	// Forma 1
	var puntero1 *int
	var variable int = 5
	fmt.Println("puntero1", puntero1) // nil
	puntero1 = &variable              // asigna la referencia de memoria de variable
	fmt.Println("puntero1 con referencia de memoria asignada", puntero1)
	fmt.Printf("El puntero 1 tiene como direcciòn %v = %d \n", puntero1, *puntero1)

	// Forma 2
	puntero2 := new(int) // reserva un espacio de memoria con el tipo de dato
	fmt.Println("puntero2", puntero2)
	fmt.Printf("El puntero 2 tiene como direcciòn %v = %d", puntero2, *puntero2)
}
