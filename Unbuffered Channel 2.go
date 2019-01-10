package main

import (
	"fmt"
	"time"
)
func Generarnumero(out chan <- int) {
	for x := 0; x < 5 ; x++  {
		out <- x
	}
	close(out)

}

func Elevarcuadrado(in <- chan int, out chan <- int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func Imprimirnumeros (in chan int ){
	for x := range in{
		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	numero := make(chan int)
	cuadrado := make(chan int)

	//Se generan los numeros

	go Generarnumero(numero)

	// Elevamos al cuadrado

	go Elevarcuadrado(numero,cuadrado)

	//Imprimimos en main

	Imprimirnumeros(cuadrado)

}
