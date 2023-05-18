package funciones

import (
	"fmt"
)

func Exponencia(valor int) {
	if valor > 100000000 {
		return
	}

	fmt.Println(valor)
	Exponencia(valor * 2) // cuando el valor pasa de 100000000 la funcion termina entrra en el return y termina y despues de eso ys va resolviendo las siguientes llamadas porque después de esta linea no tenemos ninguna instrucción más
}
