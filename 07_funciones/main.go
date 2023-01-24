package main

import "fmt"

/* Una funciòn es un conjunto de sentencias que se encarga de procesar unas entradas para generar una salida
- se declaran usando la palabra reservada func
*/

func imprimir(nombre string, apellido string, edad int) (string, int) {
	// Es posible retornar màs de una salida definiendolos en ()
	return "Hola " + nombre + " " + apellido, edad
}

func main() {
	// saludo, _ := imprimir("Edgar", "de la Cruz", 32)
	// con _ podemos omitir el valor de un paràmetro
	saludo, edad := imprimir("Edgar", "de la Cruz", 32)
	fmt.Println(saludo)
	fmt.Println("Tu edad es:", edad)

	lugarNacimiento := "Tabasco"

	// funciones anonimas son aquellas que no tienen nombre
	saludos, edad := func(nombre string, lugarNacimiento string, edad int) (string, int) {
		return "Hola me llamo " + nombre + " " + "saludos desde " + lugarNacimiento, edad
	}("Edgar de la Cruz", lugarNacimiento, 32) // pasar paràmetros

	fmt.Println(saludos, edad)

	// variables funciones: un variable almacena un tipo de datos funciòn

	f := func(a float32, b float32) float32 {
		var resultado float32 = (b * a) / 2
		return resultado
	}

	resultado := f(4, 3)
	fmt.Println(resultado)

}
