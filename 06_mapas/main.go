package main

import "fmt"

// Un mapa es una colecciòn de datos que implementa una llave-valor para integrar datos en su interior
func main() {
	// declaraciòn
	var meses map[int]string // [tipo dato clave]tipo dato valor
	// inicializar el mapa con make para crea la estructura en memoria
	meses = make(map[int]string)

	// asignamos valor
	meses[1] = "Enero"
	meses[2] = "Febrero"
	meses[3] = "Marzo"
	meses[4] = "Abril"
	meses[5] = "Mayo"
	meses[6] = "Junio"
	meses[7] = "Julio"
	meses[8] = "Agosto"
	meses[9] = "Septiembre"
	meses[10] = "Octubre"
	meses[11] = "Noviembre"
	meses[12] = "Diciembre"

	fmt.Println(meses)
	fmt.Printf("El tamaño del mapa es %d \n", len(meses))

	delete(meses, 8) // eliminar
	fmt.Printf("El tamaño del mapa ahora es %d \n", len(meses))
	fmt.Println(meses)

	// sobreescribir
	meses[12] = "Navidad"

	fmt.Println(meses)

	// obtener y validar un elemento dentro del mapa
	elemento, encontrado := meses[12] // las funciones y ciertas condiciones pueden retornar varias variables
	fmt.Printf("Elemento encontrado? %v : elemento fue %v \n", encontrado, elemento)
}
