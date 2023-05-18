package main

import (
	"github.com/damieva/go-course/middleware"
)

func main() {
	//fmt.Println("Hola Mundo!!")
	//variables.MuestroEnteros()
	//variables.RestoVariables()
	//estado, texto := variables.ConviertoATexto(1588)
	//fmt.Println(estado)
	//fmt.Println(texto)

	/*
		if os := runtime.GOOS; os == "linux" || os == "OS X." {
			fmt.Println("Esto no es Windows, es ", os)
		} else {
			fmt.Println("Esto es Windows")
		}

		switch os := runtime.GOOS; os {
		case "linux":
			fmt.Println("Esto es Linux")
		case "darwin":
			fmt.Println("Esto es Darwin")
		default:
			fmt.Printf("%s \n", os)
		}
	*/

	/*
		numero, texto := ejercicios.Ejercicio01("500")
		fmt.Println(numero)
		fmt.Println(texto)
	*/

	//Teaclado
	//teclado.IngresoNumeros()

	//Iteraciones
	//iteraciones.Iterar()

	//Ejercicio2
	//fmt.Println(ejercicios.Ejercicio2())

	//Files
	//files.GrabaTabla()
	//files.SumaTabla()
	//files.LeoArchivo()
	//files.LeoArchivo2()

	//Funciones
	//funciones.Calculos()
	//funciones.LlamarClousure()

	//Recursion
	//funciones.Exponencia(2)

	//Arreglos y Slices
	//arreglos_slices.MuestroArreglos()
	//arreglos_slices.MuestroSlice()
	//arreglos_slices.Capacidad()

	//Mapas
	//mapas.MostrarMpaas()

	//Esturcturas
	//users.AltaUsuario()

	//Interfaces
	// En ejer_interfaces.HumanosRespirando se pide como parámetro un humano, pero como modelos.Hombre es una estructura que implementa todos los metodos de Humano la toma como válida
	//Pedro := new(m.Hombre)
	//Maria := new(m.Mujer)
	//e.HumanosRespirando(Pedro)
	//e.HumanosRespirando(Maria)
	//e.EstaVivo(Pedro)
	//e.EstaVivo(Maria)

	//Defer y Panic
	//defer_panic.VemosDefer()
	//defer_panic.EjemploPanic()

	//GoRoutines
	//canal1 := make(chan bool)
	//go goroutines.MiNombreLento("Denis Amieva", canal1) //con esto lanzamos una funcion de forma asincrona(no bloqueante)
	//<-canal1 //Espera a q el canal1 termine la ejecucion, es una bandera. Tambien podemos traernos de canal1 una variable referente al resultado de la funcion
	//fmt.Println("Estoy aqui")
	//var x string
	//fmt.Scanln(&x)

	//WebServer:
	//webserver.MiWebServer()

	//Middleware:
	middleware.MiMiddleware()
}
