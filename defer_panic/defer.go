package defer_panic

import (
	"fmt"
	"log"
)

// El Defer nos permitira especificar cual sera la ultima instruccion a ejecutar antes de salir de la funcion

func VemosDefer() {
	fmt.Println("Este es un primer mensaje")
	defer fmt.Println("Este es el mensaje final")
	fmt.Println("Este es un segundo mensaje")
}

func EjemploPanic() {
	defer func() {
		reco := recover() // El recover siempre se usa con Defer para capturar el flujo despues del panic
		if reco != nil {
			log.Fatalf("Ocurrio un error que genero Panic \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("Se encontro el valor 1") //Aborta la ejecucion con un mensaje
	}
}
