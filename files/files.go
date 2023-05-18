package files

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/damieva/go-course/ejercicios"
)

var fileName string = "./files/txt/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.Ejercicio2()
	archivo, err := os.Create("./files/txt/tabla.txt")
	if err != nil {
		fmt.Println("Error al crear el archivo " + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicios.Ejercicio2()
	if Append(texto) == false {
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(texto string) bool {
	arch, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644) //con el operador '|' se envian 2 parametros, elprimero es para uq abra el archivo en modo escritura y el segundo que lo abra en modo append (xq si no machacaría el contenido)
	if err != nil {
		fmt.Println("Error durante el OpenFile " + err.Error())
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el WriteString " + err.Error())
		return false
	}

	arch.Close()
	return true
}

func LeoArchivo() {
	archivo, err := ioutil.ReadFile(fileName) //ioutil esta deprecated
	if err != nil {
		fmt.Println("Error al leer el archivo " + err.Error())
		return
	}
	fmt.Println(string(archivo))
}

func LeoArchivo2() {
	archivo, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error al abrir el archivo " + err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() { //Scan lee una linea y devuelve un bool si puede leer esa linea
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}
	archivo.Close()
}
