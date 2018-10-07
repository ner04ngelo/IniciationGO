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

	/*maps := make(map[string]string)
	fmt.Println(maps)

	mapExample := make(map[string]string, 2)
	fmt.Println(mapExample)

	maps["Nombre"] = "Paola"
	maps["Edad"] = "20"

	fmt.Println(maps["Nombre"])
	fmt.Println(maps["Edad"])*/

	/*fmt.Println(printExample())*/

	/*multi := multiplicar

	fmt.Println(multi(1,2,3,4))*/

	/*///Ejemplo de closure
	result := 0
	numbers := [...]int{
		1,
		2,
		4,
		5,
		6,
	}

	suma := func() {
		for _, numero := range numbers {
			result += numero
		}

		fmt.Printf("La suma es : %d \n", result)
	}

	suma()

	if suma != nil {
		fmt.Println("Suma tiene implementación")
	}*/

	/*inc := incremementar()

	fmt.Println("Valor de i: ", inc())
	fmt.Println("Valor de i: ", inc())
	fmt.Println("Valor de i: ", inc())
	fmt.Println("Valor de i: ", inc())*/

	n, v := retornoMultiple()

	fmt.Printf("Número: %d , Cadena: %q ", n, v)

}

func retornoMultiple() (numero int, cadena string) {
	numero = 1
	cadena = "Hola"

	return

}

/*func incremementar() func() int {
	i := 0
	return func() (r int) {
		r = i
		i += 2
		return
	}
}*/

/*func multiplicar(numbers ...int8) int {
	result := 1

	for _, number := range numbers {
		result *= int(number)
	}

	return result
}*/

/*func printExample() (name string) {
	name = "Jasson"
	return

}*/
