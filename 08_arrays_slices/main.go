package main

import "fmt"

/*
Los arrays nos permite almacenar un serie de datos y tienen una logitud fija
Un slice es un array pero con tamaño variable o dinàmico
*/
func main() {
	// arrays
	arregloColores := [3]string{"Amarrillo", "Azul", "Rojo"} // 3 posiciones en memoria
	fmt.Println(arregloColores)

	// slices
	sliceColores := []string{"Amarillo"}
	fmt.Println(sliceColores)
	// agregar elementos append()
	sliceColores = append(sliceColores, "Verde")
	fmt.Println(sliceColores)
	sliceColores = append(sliceColores, "Naranja", "Cafè")
	fmt.Println(sliceColores)
	sliceNuevosColores := []string{"Rosa", "Beige"}
	sliceColores = append(sliceColores, sliceNuevosColores...)
	fmt.Println(sliceColores)

	// cap() - capacidad slice
	sliceNumeros := []int{1, 2, 3} // 3 posiciones reservadas en memoria
	// go duplica el espacio de memoria de acuerdo a nùmero de elementos inicializados para no tener que reservar durante ejecuciòn
	sliceNumeros = append(sliceNumeros, 4)
	sliceNumeros = append(sliceNumeros, 5)
	sliceNumeros = append(sliceNumeros, 6)
	sliceNumeros = append(sliceNumeros, 7)
	fmt.Println("El tamaño del slice es:", len(sliceNumeros))
	fmt.Println("La capacidad del slice es:", cap(sliceNumeros))

	for index, elemento := range sliceNumeros {
		fmt.Println("Index:", index, "Elemento:", elemento)
	}
}
