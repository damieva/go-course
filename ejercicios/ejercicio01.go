package ejercicios

import (
	"strconv"
)

func Ejercicio01(texto string) (int, string) {
	//entero, _ := strconv.Atoi(texto) // con _ indicamos que cierto valor no nos interesa (en este caso el parametro devuelto 'error')
	entero, err := strconv.Atoi(texto)
	if err != nil {
		return 0, "Hubo un error " + err.Error()
	} else {
		if entero > 100 {
			return entero, "Es mayor que 100"
		} else {
			return entero, "Es menor que 100"
		}
	}

}
