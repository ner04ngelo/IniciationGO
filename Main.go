package main

import (
	"fmt"
	"strconv"
)

func main() {

///String


var stringExample string

stringExample= "Edad : "
edad:=30

fmt.Println(stringExample + strconv.Itoa(edad))

i:=5

	if i <= edad {
		fmt.Println("Usted es menor de edad")
	} else{
		fmt.Println("Usted tiene 30 años")
	}

}
