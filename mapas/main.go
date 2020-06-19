package main

import "fmt"

func main() {
	//se le define el 5 para darle un espacio inicial por si incrementa go lo maneja de mejor forma
	paises := make(map[string]string, 5)

	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises)

	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30,
	}

	//agregando valores al mapa
	campeonato["River Plate"] = 25
	//modificar elemento del mapa
	campeonato["Chivas"] = 55
	//borraar elementos del mapa
	delete(campeonato, "Real Madrid")

	fmt.Println(campeonato)

	//recorrer el mapa

	for equipo, puntaje := range campeonato {
		fmt.Printf("El equipo: %s, puntaje: %d \n", equipo, puntaje)
	}

	//validar si existe un valor en el mapa

	puntaje, existe := campeonato["Mineiro"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe %t \n", puntaje, existe)
}
