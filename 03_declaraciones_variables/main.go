package main

import "fmt"

var ( // puedes declarar un bloque de variables al nivel del nombre del paquete
	nombreCurso = "2022 - APRENDE A PROGRAMAR CON GO DESDE 0"
	versionGo   = "1.19"
)

func main() {
	// declaraciòn implìcita sin definirlas con var
	numero := 5
	flotante := 3.5
	cadena := "cadena"
	boleano := true

	// declaraciòn explìcita
	var numero2 int
	var flotante2 float32
	var cadena2 string
	var boleano2 bool

	fmt.Println("Declaraciòn implìcita:")
	fmt.Println(numero)
	fmt.Println(flotante)
	fmt.Println(cadena)
	fmt.Println(boleano)

	fmt.Println("Declaraciòn explìcita:")
	fmt.Println(numero2)
	fmt.Println(flotante2)
	fmt.Println(cadena2)
	fmt.Println(boleano2)

	fmt.Println(nombreCurso)
	fmt.Println(versionGo)
}
