package main

import "fmt"
import "./File"
/// Crear un propio tipo
type Money int

///////////////
func printExample() {
	fmt.Println("Mundo")
}

func pointerExample(c *int) {
	*c += 100
}

func (d Money) String() string {
	return fmt.Sprintf("$%d", d)
}

/// Crear una structura
/*type Persona struct {
	Nombre string
	edad   int
}*/



func main() {

	employee := File.Empleado{
		Persona: File.Persona{
			Nombre: "Manolo",
			Edad: 40,
		},
		Puesto: "Contador",
	}

	employee.SobreEscribirPuesto("Vendedor")

	fmt.Println(employee.ImprimirPuesto())
	fmt.Println(employee.Edad)
	fmt.Println(employee.Puesto)

	/*var p Persona
	p.Nombre = "Jasson"
	p.edad = 21*/

	/*///Punteros

	a := 25
	fmt.Println("Valor de a:", a)
	fmt.Println("Dirección de memoria de a:", &a)

	b := &a

	fmt.Println("Valor de b:", b)
	fmt.Println("B contiene: ", *b)



	pointerExample(b)

	fmt.Println("Valor de a:", a)
	fmt.Println("B contiene: ", *b)*/

	//defer multiplicar(1, 2, 5)
	//printName()

	//// Leer archivos

	/*f, err := os.Open("/Users/USER/Desktop/texto.txt")

	if err != nil {
		panic(err)
	}

	data := make([]byte, 20)

	c, err := f.Read(data)

	if err != nil {
		panic(err)
	}

	fmt.Printf(" Cantidad de bytes leidos: %d\n Texto leido: %q\n error: %v", c, data, err)

	f.Close()*/
	//defer printExample()
	//fmt.Println("Hola")
	//
	//defer func() {
	//	cadena := recover()
	//	fmt.Println(cadena)
	//}()
	//panic("Error")

}

/*func multiplicar(numbers ...int8) {
	result := 1

	for _, number := range numbers {
		result *= int(number)
	}

	fmt.Println(result)
}*/

/*func printName() {
	fmt.Println("Ejemplo de defer")
}*/

//func main() {
//
//	//sliceExample := make([]byte, 4 , 10)
//	//
//	//sliceExample = []byte{'H','O','L','A'}
//	//
//	//fmt.Println(sliceExample)
//	//fmt.Printf("SliceExample: %q \n", sliceExample)*/
//	//
//	///*var y []int
//	//
//	//for i := 1; i < 13; i++ {
//	//	y = append(y, i)
//	//	fmt.Println("Array")
//	//	fmt.Println(y)
//	//	fmt.Printf("Longitud y: %d, Capacidad y: %d, Elementos y: %d \n", len(y), cap(y), i)
//	//}*/
//	//
//	///*	sliceOrigin := []int{1, 2, 3}
//	//	sliceDestiny := []int{4, 5}
//	//	copy(sliceDestiny, sliceOrigin)
//	//	fmt.Println(sliceOrigin, sliceDestiny)*/
//	///*
//	//	objects := []string{
//	//		"Mesa",
//	//		"Silla",
//	//		"Ventilador",
//	//		"Licuadora",
//	//	}
//	//
//	//
//	//	for _, obj := range objects{
//	//		fmt.Printf("El objeto es %q \n", obj )
//	//	}*/
//	//
//	///*maps := make(map[string]string)
//	//fmt.Println(maps)
//	//
//	//dias := map[int]string{
//	// 1: "Lunes",
//	// 2: "Martes",
//	// 3: "Miercoles",
//	//}
//	//
//	//fmt.Println(dias)
//	//
//	//mapExample := make(map[string]string, 2)
//	//fmt.Println(mapExample)
//	//
//	//maps["Nombre"] = "Paola"
//	//maps["Edad"] = "20"
//	//
//	//fmt.Println(maps["Nombre"])
//	//fmt.Println(maps["Edad"])*/
//	//
//	///*fmt.Println(printExample())*/
//	//
//	///*multi := multiplicar
//	//
//	//fmt.Println(multi(1,2,3,4))*/
//	//
//	///*///Ejemplo de closure
//	//result := 0
//	//numbers := [...]int{
//	//	1,
//	//	2,
//	//	4,
//	//	5,
//	//	6,
//	//}
//	//
//	//suma := func() {
//	//	for _, numero := range numbers {
//	//		result += numero
//	//	}
//	//
//	//	fmt.Printf("La suma es : %d \n", result)
//	//}
//	//
//	//suma()
//	//
//	//if suma != nil {
//	//	fmt.Println("Suma tiene implementación")
//	//}*/
//	//
//	///*inc := incremementar()
//	//
//	//fmt.Println("Valor de i: ", inc())
//	//fmt.Println("Valor de i: ", inc())
//	//fmt.Println("Valor de i: ", inc())
//	//fmt.Println("Valor de i: ", inc())*/
//	//
//	///*n, v := retornoMultiple()
//	//
//	//fmt.Printf("Número: %d , Cadena: %q ", n, v)
//
//}

/*func retornoMultiple() (numero int, cadena string) {
	numero = 1
	cadena = "Hola"

	return

}*/

/*func incremementar() func() int {
	i := 0
	return func() (r int) {
		r = i
		i += 2
		return
	}
}*/

/*func printExample() (name string) {
	name = "Jasson"
	return

}*/
