package main

import (
	"fmt"
	"time"
)

///Ejemplo de Unbuffered Channel
func main() {
	numero := make(chan int)
	cuadrado := make(chan int)

	//Se generan los numeros

	go func() {
		for x := 0; x < 5 ; x++  {
			numero <- x
		}
		close(numero)

	}()

	// Elevamos al cuadrado

	go func() {
		for x := range numero {
			cuadrado <- x * x
		}
		close(cuadrado)
	}()

	//Imprimimos en main

	for x := range cuadrado{
		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}
}
