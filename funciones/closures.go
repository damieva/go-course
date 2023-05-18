package funciones

import (
	"fmt"
)

func tabla(valor int) func() int { // la funcion privada 'tabla' devuelve a su vez una funcion anonima la cual devuelve un int
	numero := valor
	secuencia := 0 // los clousures nos sirven para preservar y ofuscar los valores tanto de numero como de secuencia, este ultimo se ira modificando segun vayamos llamando n veces a la funcion (esto será como mantener una foto exacta de cada ejecución de la funcion en la vble de abajo llamada 'MiTabla')
	return func() int {
		secuencia++
		return numero * secuencia
	}
}

func LlamarClousure() {
	tabladel := 2
	MiTabla := tabla(tabladel)
	for i := 1; i <= 10; i++ {
		fmt.Println(MiTabla())
	}
}
