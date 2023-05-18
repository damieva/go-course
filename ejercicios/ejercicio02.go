package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Ejercicio2() string {
	scanner := bufio.NewScanner(os.Stdin)
	var num int
	var err error
	var retorno string

	for {
		fmt.Println("Introduce un numero: ")
		if scanner.Scan() {
			num, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		retorno += fmt.Sprintf("%d \n", num*i)
	}

	return retorno
}
