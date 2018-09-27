package main

import "fmt"

func main() {

	var choice int8

	fmt.Println("Ingresa un número")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		fmt.Println("Hoy lloverá")
	case 2:
		fmt.Println("No va a llover")

	default:
		fmt.Println("Opción no disponible")
	}

}
