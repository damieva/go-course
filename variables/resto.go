package variables // todo lo q este dentro del package variables podra ser utilizado sin problema, ejemplo: funcón MuestraVriables

import (
	"fmt"
	"strconv"
	"time"
)

// tanto vbles como funciones que empiecen en mayuscula podrán ser vista por resto de archivos .go dentro del package o archivos .go que importen el package

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	Nombre = "Pedro"
	Estado = true
	Sueldo = 1577.66
	Fecha = time.Now()

	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}

func ConviertoATexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero) //con el : cogemos el tipo de dato del valor que le asignamos
	return true, texto
}
