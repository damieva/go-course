package mapas

import (
	"fmt"
)

func MostrarMpaas() {
	paises := make(map[string]string) //map[clave]valor
	//fmt.Println(paises)

	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises)
	fmt.Println(paises["Argentina"])

	campeonato := map[string]int{
		"Barcelona":   39,
		"Real Madrid": 38,
		"Chivas":      37,
		"Boca":        30,
	}

	fmt.Println(campeonato) //ordena por clave no por valor

	for equipo, puntos := range campeonato {
		// 'for... range' es similiar al 'for each'
		// al iterar nunca ordena
		fmt.Printf("Equipo %s, tiene %d puntos \n", equipo, puntos)
	}

	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["Juventus"]
	fmt.Printf("Los puntos capturados son %d y el equipo existe = %t \n", puntaje, existe)
}
