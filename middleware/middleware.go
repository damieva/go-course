package middleware

import (
	"fmt"
)

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func MiMiddleware() {
	fmt.Println("Inicio")
	result := operacionesMidd(sumar)(2, 3) //operacionesMidd recibe la funcion sumar como parámetro para después configurar su función (de retorno) a partir de la funcion recibida y alguna instruccion mas. Finalmente esa funcion devuelta por el Middleware toma (2,3) como parametros y hace la operacion.
	fmt.Println(result)
	result = operacionesMidd(restar)(6, 3)
	fmt.Println(result)
	result = operacionesMidd(multiplicar)(2, 4)
	fmt.Println(result)
}

// La funcion de middleware toma la funcion que recibe como parametro y devuelve otra funcion en la que utiliza esa misma funcion recibida como parametro
func operacionesMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de operacion")
		return f(a, b)
	}
}
