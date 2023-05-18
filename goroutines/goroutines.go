package goroutines

import (
	"fmt"
	"strings"
	"time"
)

func MiNombreLento(nombre string, canal1 chan bool) {
	letras := strings.Split(nombre, "") //al poner la cadena vacía va a coger cada letra de nombre y va a conformar con ella un elemento del vector
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
	// al finalizar el bucle vamos a asignar un valor a ese canal
	canal1 <- true
}
