package main

import "fmt"

func main() {
	paises := make(map[string]string, 5) //make crea un slice de lenght 0 y size 5. El slice puede crecer dinamicamente pero es más lento q reservando el espacio cuando se crea
	fmt.Println(paises)

	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises["Mexico"])
	fmt.Println(paises)

	campeoneto := map[string]int{ //Por defecto ordena alfabeticamente por la clave
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30}

	campeoneto["River Plate"] = 25    // Añadimos un nuevo elemento
	campeoneto["Chivas"] = 55         // Modificamos un elemento existente
	delete(campeoneto, "Real Madrid") // Eliminamos un elemento
	fmt.Println(campeoneto)

	for equipo, puntos := range campeoneto { // Recorrer mapa
		fmt.Printf("El equipo %s, tiene %d puntos \n", equipo, puntos)
	}

	puntos, existe := campeoneto["Atletico de Madrid"] // Check de si elemento existe
	fmt.Printf("El puntaje capturado es %d, y el equipo existe %t \n", puntos, existe)

	puntos, existe = campeoneto["Chivas"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe %t \n", puntos, existe)
}
