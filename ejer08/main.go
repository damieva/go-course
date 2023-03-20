package main

import "fmt"

var tabla [10]int

func main() {
	//VECTORES
	tabla[0] = 1
	tabla[5] = 15
	fmt.Println(tabla)

	tabla := [10]int{1, 2, 3, 5}
	fmt.Println(tabla)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	//MATRICES
	var matriz [5][7]int
	matriz[3][5] = 1
	fmt.Println(matriz)

	//SLICES: vectores dinamicos (pueden cambiar su tamaño despues de haberlo creado)
	slice := []int{2, 3, 4} //no se indica tamaño
	fmt.Println(slice)
	//podemos crear un slice en base a una porcion de otro vector
	elementos := [5]int{1, 5, 9, 8, 7}
	porcion := elementos[2:5]
	fmt.Println(porcion)

	elementos_2 := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos_2), cap(elementos_2))

	elementos_3 := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		elementos_3 = append(elementos_3, i)
	}
	//cuando la capacidad del vector es superada al añadir elementos, Go reserva capacidad automaticamente en base a potencias de 2
	//2,4,8,16,32,64,128,256,512,1024,...
	fmt.Printf("Largo %d, Capacidad %d", len(elementos_3), cap(elementos_3))
}
