package arreglos_slices

import (
	"fmt"
)

var tabla [10]int = [10]int{10, 5, 8} // Inicializacmos el array, el resto de posiciones que no introducimos las deja a 0
var matriz [20][30]int

func MuestroArreglos() {
	tabla[7] = 33
	tabla[2] = 54
	fmt.Println(tabla)

	tabla2 := [10]string{"pablo", "juan"}
	fmt.Println(tabla2)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matriz[7][24] = 15
	fmt.Println(matriz)
}
