package main


//// wg es utilizado para indicarle al programa que debe esperar
//// que finalicen las gorutinas
//var wg sync.WaitGroup
//
//func main() {
//
//
//	// Añadimos 2 a wg para que espero que finalicen 2 gorutinas.
//	wg.Add(2)
//	fmt.Println("Iniciamos las gorutinas...")
//
//	go imprimirCantidad("A")
//
//	go imprimirCantidad("B")
//	// Esperamos a que las gorutinas finalicen.
//	fmt.Println("Esperando que Finalicen...")
//	wg.Wait()
//	fmt.Println("\nTerminando el programa")
//
//	/*	go numbers()
//		go alphabets()
//		time.Sleep(3000 * time.Millisecond)
//		fmt.Println("main terminated")*/
//
//
//}
//
//
//func imprimirCantidad(etiqueta string) {
//	// Llamamos la funcion Done() de wg para indicale que la gorutina termino.
//	defer wg.Done()
//	// Espera aleatoria
//	for cantidad := 1; cantidad <= 10; cantidad++ {
//		sleep := rand.Int63n(1000)
//		time.Sleep(time.Duration(sleep) * time.Millisecond)
//		fmt.Printf("Cantidad: %d de %s\n", cantidad, etiqueta)
//
//	}
//}
//
//
///*func numbers() {
//	for i := 1; i <= 5; i++ {
//		time.Sleep(250 * time.Millisecond)
//		fmt.Printf("%d ", i)
//	}
//}
//func alphabets() {
//	for i := 'a'; i <= 'e'; i++ {
//		time.Sleep(400 * time.Millisecond)
//		fmt.Printf("%c ", i)
//	}
//}