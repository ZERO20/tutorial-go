package main

import (
	"fmt"
	"time"
)

/*
para declararlas
- go funcion()

- procesan al mismo tiempo
*/
func main() {
	nombres := []string{"Edgar", "Paula", "Gaby", "Brandy", "Piter"}
	go Saludar(nombres)
	go Despedir(nombres)

	//go fmt.Println("Hola mundo")
	teclado := ""
	fmt.Scan(&teclado) // permite input de datos del teclado
}

func Saludar(nombres []string) {
	for _, nombre := range nombres {
		fmt.Printf("Hola %s \n", nombre)
		time.Sleep(time.Second)
	}
}

func Despedir(nombres []string) {
	for _, nombre := range nombres {
		fmt.Printf("Adiós %s \n", nombre)
		time.Sleep(time.Second / 2)
	}
}
