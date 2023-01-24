package main

import "fmt"

func main() {
	// if / else
	edad := 70
	if edad < 18 {
		fmt.Println("Eres menor de edad")
	}
	// con declaraciòn de variable en misma lìnea
	if genero := 'F'; genero == 'M' {
		fmt.Println("Gènero Masculino")
	} else {
		fmt.Println("Gènero Femenino")
	}

	// else if
	if edad < 18 {
		fmt.Println("Eres menor de edad")
	} else if edad < 50 {
		fmt.Println("Eres un adulto")
	} else {
		fmt.Println("Eres una persona de la tercer edad")
	}

	// switch
	// Validar con valores en especifico

	// Especificando la expresiòn al inidio
	switch edad {
	case 10:
		fmt.Println("La edad es 10")
	case 20:
		fmt.Println("La edad es 20")
	case 30:
		fmt.Println("La edad es 30")
	case 40:
		fmt.Println("La edad es 40")
	default:
		fmt.Println("No se encontrò la edad")
	}

	// Sin especificar la expresiòn y evaluarla en los cases
	switch {
	case edad < 10:
		fmt.Println("La edad es 10")
	case edad < 20:
		fmt.Println("La edad es 20")
	case edad < 30:
		fmt.Println("La edad es 30")
	case edad < 40:
		fmt.Println("La edad es 40")
	case edad < 70:
		fmt.Println("La edad es 70")
	default:
		fmt.Println("No se encontrò la edad")
	}
}
