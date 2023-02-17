package main

import (
	"fmt"
	"strconv"
)

// data types
var numero int //float32 float64 int8 int16
var texto string
var status bool = true

func main() {
	numero2, numero3, numero4, texto, status := 2, 5, 76, "texto", false // con el := se crea una nueva vble

	var moneda float32 = 0

	numero2 = int(moneda) //casting de float a entero

	texto = fmt.Sprintf("%f", moneda) //casting de float32 a string
	fmt.Println(texto)

	texto = strconv.Itoa(int(moneda))
	fmt.Println(texto)

	fmt.Println(numero)
	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(status)

	MostrarStatus() // con la mayuscula vemos esta funcion desde otros paquetes

}

func MostrarStatus() {
	fmt.Println(status)
}
