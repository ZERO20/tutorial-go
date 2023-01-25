package main

import (
	"fmt"
	"math"
)

func main() {
	numero := -9.0
	numeroAbs := math.Abs(numero)              // número absoluto
	numeroRaizCuadrada := math.Sqrt(numeroAbs) //raíz cuadrada

	fmt.Println("Número: ", numero)
	fmt.Println("Valor absoluto: ", numeroAbs)
	fmt.Println("Raíz cuadrada: ", numeroRaizCuadrada)

	// floor: redondea el número hacía abajo
	numero = 3.5
	numero = math.Floor(numero)
	fmt.Println(numero)

	// ceil: redondea el número hacía arriba
	numero = 3.5
	numero = math.Ceil(numero)
	fmt.Println(numero)

	//round: redondeo
	numero = 3.4 // 3.5  -> 4 y 3.4 -> 3
	numero = math.Round(numero)
	fmt.Println("El número redondeado es: ", numero)

	//Base elevar al exponente
	numero = 2
	numero = math.Pow(numero, 2)
	fmt.Println("El exponente es: ", numero)

	//Módulo
	numero = 10
	numero = math.Mod(numero, 2)
	fmt.Println("El módulo es: ", numero)

	//Truncar: eliminar parte decimal
	numero = 10.583422
	numero = math.Trunc(numero)
	fmt.Println("El número truncado es: ", numero)

	//Calcular número mayor
	numero1 := 3.4
	numero2 := 4.4
	numero = math.Max(numero1, numero2)
	fmt.Println("El número mayor es: ", numero)

	//Calcular número menor
	numero = math.Min(numero1, numero2)
	fmt.Println("El número menor es: ", numero)
}
