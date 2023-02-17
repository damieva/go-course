package main

import (
	"bufio"
	"fmt"
	"os"
)

var numero1 int
var numero2 int
var resultado int
var leyenda string

func main() {
	fmt.Println("Ingrese num 1:")
	fmt.Scanf("%d", &numero1)

	fmt.Println("Ingrese num 2:")
	fmt.Scanf("%d", &numero2)

	fmt.Println("Descripcion: ")

	scanner := bufio.NewScanner(os.Stdin) // creo un objeto Scanner que me va a leer de teclado
	if scanner.Scan() {                   // se no ingresamos nada por teclado ahí no entrará, es un try-catch
		leyenda = scanner.Text()
	}

	resultado = numero1 + numero2
	fmt.Println(leyenda, resultado)
}
