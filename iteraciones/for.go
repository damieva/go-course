package iteraciones

import "fmt"

func Iterar() {
	for i := 0; i < 10; i++ {
		if i == 8 {
			break // sale del bucle
		}

		if i == 6 {
			continue //vuelve al principio del bucle
		}
		fmt.Println(i)
	}
}
