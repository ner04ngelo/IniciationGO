package main

import "fmt"

func main() {

	/*sliceExample := make([]byte, 4 , 10)

	sliceExample = []byte{'H','O','L','A'}

	fmt.Println(sliceExample)
	fmt.Printf("SliceExample: %q \n", sliceExample)*/

	/*var y []int

	for i := 1; i < 13; i++ {
		y = append(y, i)
		fmt.Println("Array")
		fmt.Println(y)
		fmt.Printf("Longitud y: %d, Capacidad y: %d, Elementos y: %d \n", len(y), cap(y), i)
	}*/

	/*	sliceOrigin := []int{1, 2, 3}
		sliceDestiny := []int{4, 5}
		copy(sliceDestiny, sliceOrigin)
		fmt.Println(sliceOrigin, sliceDestiny)*/
	/*
		objects := []string{
			"Mesa",
			"Silla",
			"Ventilador",
			"Licuadora",
		}


		for _, obj := range objects{
			fmt.Printf("El objeto es %q \n", obj )
		}*/

	maps := make(map[string]string)
	fmt.Println(maps)

	mapExample := make(map[string]string, 2)
	fmt.Println(mapExample)

	maps["Nombre"] = "Paola"
	maps["Edad"] = "20"

	fmt.Println(maps["Nombre"])
	fmt.Println(maps["Edad"])
}
