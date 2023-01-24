package main

import (
	"fmt"
	"sync"
	"time"
)

/*
- WaitGroup
  - nos permite sincronizar nuestras gorutinas con el thread principal de la aplicación
  - paquetes sync
  - WaitGroup.Add(cantidad) cuenta cuantas gorutinas vamos a ejecutar
  - WaitGroup.Done() informamos que la gorutina terminó
  - WaitGroup.Wait() le indica al hilo principal que debe esperar que todas las gorutinas terminen o llamen al Done del WaitGroup dentro
    de la función

- Mutex
  - Se implementan para trabajar con concurrencia, al tratar de acceder más de una vez al mismo recurso
  - resuelve problemas de Data Race (acceso al mismo recurso)
  - Paquete sync
  - Mutex.Lock()
  - Mutex.Unlock()
*/
func main() {
	var wg sync.WaitGroup
	wg.Add(2)       // cantidad de gorutinas
	defer wg.Wait() // se puede utilizar defer

	nombres := []string{"Edgar", "Paula", "Gaby", "Brandy", "Piter"}

	//wg.Wait() //espera la ejecución de los done de cada gorutina
	// Mutext
	var nombreTemporal string = ""
	var mx sync.Mutex

	go Saludar(nombres, &wg, &nombreTemporal, &mx)
	go Despedir(nombres, &wg, &nombreTemporal, &mx)
}

func Saludar(nombres []string, wg *sync.WaitGroup, nombreTemporal *string, mx *sync.Mutex) {
	defer wg.Done() //utilizando defer
	for _, nombre := range nombres {
		//fmt.Printf("Hola %s \n", nombre)
		// mutex
		mx.Lock()
		nombreTemporal = &nombre
		fmt.Printf("Hola %s \n", *nombreTemporal)
		time.Sleep(time.Second)
		mx.Unlock()
	}
	//wg.Done() // terminó la ejecución
}

func Despedir(nombres []string, wg *sync.WaitGroup, nombreTemporal *string, mx *sync.Mutex) {
	defer wg.Done() //utilizando defer
	for _, nombre := range nombres {
		//fmt.Printf("Adiós %s \n", nombre)
		// time.Sleep(time.Second / 2)
		// mutex
		mx.Lock()
		nombreTemporal = &nombre
		fmt.Printf("Adiós %s \n", *nombreTemporal)
		time.Sleep(time.Second)
		mx.Unlock()
	}
	//wg.Done() // terminó la ejecución
}
