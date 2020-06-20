package main

import (
	"log"
)

func main() {
	//Defer se ejecuta si o si cuando una funcion se va por un error, por un return o un fin de funcion
	/*archivo := "prueba.txt"
	f, err := os.Open(archivo)

	defer f.Close()

	if err != nil {
		fmt.Println("Error abriendo el archivo")
		os.Exit(1)
	}*/

	//Panic es para forzar un error y que el programa termine
	ejemploPanic()
}

func ejemploPanic() {
	//si pongo defer siempre se va a ejecutar si o si
	defer func() {
		//si en un dado caso no encuentro un panic y paso se guarda en una variable y evaluamos
		reco := recover()
		if reco != nil {
			log.Fatalf("ocurrio un error que genero Panic \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("se encontro el valor de 1")
	}
}
