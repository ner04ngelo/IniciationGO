package main

import (
	"fmt"
	"unsafe"
)

func main() {

/// Casting de variables

/// En Go el casting tiene que ser de forma explicito ya que no funciona como lenguajes de alto nivel.

var entero8 uint8   // 1 byte
var entero16 uint16 // 2 byte
var entero32 uint32 // 4 byte
var entero64 uint64 //8 byte

entero8=15
entero32=230

fmt.Println(unsafe.Sizeof(entero8), unsafe.Sizeof(entero16), unsafe.Sizeof(entero32), unsafe.Sizeof(entero64))
}




