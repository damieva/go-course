package main

import "fmt"

//FUNCIONES ANONIMAS

var Calculo func(int, int) int = func(num1 int, num2 int) int { // las funciones tb son tipos de datos
	return num1 + num2
}

func main() {
	fmt.Printf("%d\n", Calculo(5, 7))

	// al asignar una funcion a una vble podemos redefinir facilmente una funcion, solo manteniendo el prototipo de la misma
	Calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}

	fmt.Printf("%d\n", Calculo(5, 7))

	Operaciones()

	//CLOSURES: el valor de MiTabla es la funcion q va dentro de tabla
	tableDel := 2
	MiTabla := Tabla(tableDel)
	for i := 1; i < 11; i++ {
		fmt.Println(MiTabla())
	}
}

func Operaciones() {
	resultado := func() int {
		var a int = 23
		var b int = 13
		return a + b
	}
	fmt.Println(resultado())
}

func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia += 1
		return numero * secuencia
	}
}
