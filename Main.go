package main

import "fmt"

/// NOTA: En Go toda variable que se es declarada debe de ser utilizada
/// de la misma forma sucede con los import sino se usa una libreria importada el compilador obliga a eliminarla

var variableLocal = "Esto es una variable local"

func main() {
	/*
		/// La libreria fmt es una abreviación de format, esta se usa para poder imprimir e ingresar datos

		fmt.Println("Holixxx")

		/// Para declarar una variable se usa la palabra reservada VAR, haciendo referencia que es una variable
		/// después el nombre de la variable y por último el tipo
		/// a diferencia de otros lenguajes de programación la declaración en GO es al revés
		/// GO inicializa la variable con un valor por defecto

		var numero int = 10
		fmt.Println(numero)

		/// Para declarar una variable de manera rápida determinando el tipo de manera dinámica
		/// una vez que la variable se le asigna una tipo este ya no se puede cambiar
		/// el simbolo := solo se cuando se declara una nueva variable

		nombre := "hola"

		/// En GO se pueden declarar variables de forma consecutiva

		nombre, numero = "Eduardo", 20

		/// Esta forma de asignar valores a las variables de forma consecutiva puede resultar util al momento del intercambio
		/// de valores entre variables

		nombre= "Pedro"
		name := "Juan"
		nombre, name = name, nombre


		fmt.Println("Nombre: ", nombre)
		fmt.Println("Name: ", name)

		/// Esta forma de asignación se puede usar también cuando se crea una nueva variable

		nombre2, name := "Gloria", "Teresa"

		fmt.Println("Nombre2: ", nombre2)
		fmt.Println("Name: ", name)


		///Para declarar multiples variables se usa var(), dentro de los parentesis se declaran las variables con sus asignaciones

		var numero4 =70

		fmt.Println(numero4)


		var(
			variable string
			example, example2 = 1,2
		)

		fmt.Println(variable)
		fmt.Println(example)
		fmt.Println(example2)*/

		/*var simpleArray[2] int
		simpleArray[0] = 1
		simpleArray[1] = 2*/


	imprimir()

}

func imprimir() {
	fmt.Println(variableLocal)
}


/*type Student struct{
	name string
	lastName string
	age int
	average float64
	genre bool
}*/

