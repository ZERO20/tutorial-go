package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
  - un channel es un tipo de variable que permite el paso de información entre gorutinas
  - pueden ser bidireccionales o unidireccionales(<- usando antes o despúes)
  - declaración:
    ch := make(chan string)
  - Enviar y recibir
    enviar: canal <- dato
    recibir: dato := <-canal  (<-canal va pegado)
*/
func main() {
	chNumeros := make(chan int) // int tipo de dato que cruzará
	chImprimir := make(chan string)
	chPares := make(chan int)
	chImpares := make(chan int)

	// cerrar canal
	defer close(chNumeros)
	defer close(chImprimir)
	defer close(chPares)
	defer close(chImpares)

	// gorutinas
	go GenerarNumeros(chNumeros)
	go RecibirNumero(chNumeros, chPares, chImpares)
	go NumeroImpar(chImpares, chImprimir)
	go NumeroPar(chPares, chImprimir)

	// se procesa el mensaje para imprimir
	/*for {
		mensageRecibido := <-chImprimir
		fmt.Println(mensageRecibido)
	}*/

	Imprimir(chImprimir)
}

func GenerarNumeros(ch chan<- int) {
	// Gorutina que genera números
	numero := 0
	for {
		numero++
		ch <- numero            // enviar el parámetro
		time.Sleep(time.Second) //dormir 1 segundo la ejecución
	}
}

func RecibirNumero(ch <-chan int, chPar chan<- int, chImpar chan<- int) {
	// Gorutina que recibe un número y determina si es par o impar para mandarlo a su respectivo canal
	// ch <-chan int - solo lectura / ch chan<- int solo salida
	for {
		numero := <-ch
		if numero%2 == 0 {
			// Enviar al canal de pares
			chPar <- numero
		} else {
			// Enviar al canal de impares
			chImpar <- numero
		}
	}
}

func NumeroImpar(ch chan int, chImprimir chan string) {
	// Gorutina que recibe un número impar y lo manda a imprimir
	for {
		numero := <-ch
		imprimir := "Número impar = " + strconv.Itoa(numero)
		chImprimir <- imprimir
	}
}

func NumeroPar(ch chan int, chImprimir chan string) {
	// Gorutina que recibe un número par y lo manda a imprimir
	for {
		numero := <-ch
		imprimir := "Número par = " + strconv.Itoa(numero)
		chImprimir <- imprimir
	}
}

func Imprimir(ch chan string) {
	// Imprime un mensaje desde un canal
	for {
		mensageRecibido := <-ch
		fmt.Println(mensageRecibido)
	}
}
