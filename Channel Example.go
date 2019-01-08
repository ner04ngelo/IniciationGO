package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)

	go enviarPing(ch)
	go imprimirPing(ch)

	var input string
	fmt.Scanln(&input)

	fmt.Println("Fin...")
}

func enviarPing(c chan string) {

	c <- "ping"
}

func imprimirPing(c chan string) {
	var contador int

	for {
		contador++
		fmt.Println(<-c, "", contador)
		time.Sleep(time.Second * 1)
	}
}
