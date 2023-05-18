package ejer_interfaces

import (
	"fmt"

	"github.com/damieva/go-course/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando \n", hu.Sexo())
}

func EstaVivo(sv interfaces.SerVivo) {
	if sv.EstaVivo() {
		fmt.Println("Estoy vivo")
	} else {
		fmt.Println("No estoy vivo")
	}
}
