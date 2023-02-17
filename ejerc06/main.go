package main

import "fmt"

func main() {
	fmt.Println(uno(5))
	numero, estado := dos(2)
	fmt.Println(numero)
	fmt.Println(estado)
	fmt.Println(Calculo(5, 23, 45, 68))
}

func uno(numero int) (z int) { //func name(params) return
	z = numero * 2
	return
}

func dos(numero int) (int, bool) { // funcion que devolvemos varios valores
	if numero == 1 {
		return 5, false
	} else {
		return numero * 2, true
	}
}

func Calculo(numero ...int) int { // funcion que recibe un numero de parametros variables
	total := 0
	for item, num := range numero { //range toma un rango de parametros, lista o vectores. En cada iteracion devuelve 2 valores 'indice:valor'
		total += num
		fmt.Printf("%d \n", item)
	}
	return total
}
