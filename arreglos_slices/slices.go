package arreglos_slices

import (
	"fmt"
)

// El fuerte de los slices en contraposision con los arreglos es la redimension
var tablaS []int = []int{2, 5, 4} // vector sin dimension, la dimension se determina por la inicializacion de sus datos (en este caso 3)
var arreglo [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

func MuestroSlice() {
	fmt.Println(tablaS)

	porcion := arreglo[3:] // porcion es un slice desde el elemento 3 del arreglo al final
	porcion2 := arreglo[:5]
	porcion3 := arreglo[6:7]

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)

	//los slices manejan la dimension y la capacidad. Con el concepto de capacidad es como Go resuelve el problema de performance al hacer crecer este tipo de objetos.
	//	dimension: num de elementos
	//  capacidad: valor por encima de la dimension que se reserva automaticamente en memoria y permite hacer crecer rapido a los slices.

	// los arreglos son estáticos, no pueden crecer en tiempo de ejecucion. Los slices si.
}

func Capacidad() {
	// la instruccion make no permitirá crear un slice vacio y darle una capacidad.
	// la funcion make no solo es para slices si no para todo tipo de estructuras
	elementos := make([]int, 5, 20) // make(data_type, lenght, capacity)
	fmt.Printf("Largo %d, Capacidad %d \n", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 128; i++ {
		nums = append(nums, i) // voy añadiendo elementos al slice
	}
	// go automaticamente le asigna una capacidad y la va incrementando en tiempo de ejecucion a medida que el slice crece.
	// esto es una ventaja muy importante para el desarrollador porque no se tiene que preocupar de hacer crecer su slice, si no que el propio lenguaje está optimizado para ello.
	// la manera de aumentar la capacidad es: cuando la dimension sobrepasa una potencia de 2 le asigno la siguiente, ejemplo:
	// dimension 127 -> capacidad 128
	// dimension 129 -> capacidad 256
	fmt.Printf("Largo %d, Capacidad %d \n", len(nums), cap(nums))
}
