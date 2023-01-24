package main

import (
	"fmt"
	"strings"
)

/*
- Solo existe el ciclo for
- Estructura
for inicializaciòn, condiciòn, postcondiciòn {
}
- break - salir del for
- continue - skip
*/
func main() {
	// break
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 5 { // 0...5
			break
		}
	}

	fmt.Println(strings.Repeat("*", 50))
	// continue
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		} else {
			continue
		}
		fmt.Println("finalizado", i)
	}

	fmt.Println(strings.Repeat("*", 50))

	// ciclo infinito do while
	contador := 0
	for {
		fmt.Println("Hola mundo!")
		contador++
		if contador > 10 {
			break
		}
	}

	contador2 := 0
	for contador2 < 10 {
		fmt.Println("Hola amigo!")

		contador2++
	}

	// for range para iterar arrays, slices
	var items []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for index, value := range items {
		fmt.Println("index: ", index, "value: ", value)
	}
}
