package variables

import "fmt"

func MuestroEnteros() {
	var intcomun int     // automaticamente la vble tendra el valor minimo para ese tipo de dato, en este caso 0
	intde32 := int32(10) //automaticamente con := la vble toma el tipo de dato de lo que le estamos asignando
	intde64 := int64(100)

	fmt.Println("intcomun = ", intcomun)
	fmt.Println("intcomun = ", intde32)
	fmt.Println("intcomun = ", intde64)
}
