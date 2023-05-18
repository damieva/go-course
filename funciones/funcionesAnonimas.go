package funciones

// Son funciones que pueden tanto asignarse a una vble como pasarse por parámetro

import "fmt"

func Calculos() {

	var num3 int = 6
	var num4 int = 245

	calculo := func(num1 int, num2 int) int {
		return num1 + num2 + num3 + num4 // si hubiesemos creado una vble por fuera, el scope no permitiría utilizar las vbles num3 y num4
	}

	fmt.Println(calculo(2, 4))

	calculo = func(num1 int, num2 int) int { // este objecto de tipo func lo podemos redefinir las veces que queramos siempre respetando el prototipo de la funcion
		return num1 * num2 * num3 * num4
	}

	fmt.Println(calculo(2, 4))
}
