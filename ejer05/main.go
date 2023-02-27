package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	numero := 0
	for { // blucle infinito
		fmt.Println("Continuo")
		fmt.Println("Ingrese el numero secreto: ")
		fmt.Scanln(&numero)
		if numero == 100 {
			break //rompe el bucle for
		}
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("\n Valor de i: %d", i)
		if i == 5 {
			fmt.Printf(" Multiplicamos por 2 \n")
			i = i * 2
			continue // manda la ejecución al inicio del bucle for
		}
		fmt.Println("                  Pasó por aqui")
	}

	i := 0
RUTINA:
	for i < 10 {
		if i == 4 {
			i += 2
			fmt.Println("voy a RUTINA sumando 2 a i")
			goto RUTINA // a difrencia del continue, el goto va a donde le digamos
		}
		fmt.Printf("Valor de i: %d\n", i)
		i++
	}
}
