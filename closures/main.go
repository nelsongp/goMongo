package main

import "fmt"

var calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("Suma 5 + 7 = %d \n", calculo(5, 7))

	//Restamos
	calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}
	fmt.Printf("Restamos 6 - 4 = %d \n", calculo(6, 4))

	operaciones()
	tablaDel := 2
	MiTabla := tabla(tablaDel)

	for i := 1; i < 11; i++ {
		fmt.Println(MiTabla())
	}
}

func operaciones() {
	resultado := func() int {
		a := 23
		b := 23
		return a + b
	}
	fmt.Println(resultado())
}

//closures
func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	//cuando ejecutemos este valor en nuestra funcion, lo que vamos a hacer es estar ejecutando la funcion muchas veces
	//los primeros valores solo se ejecutan la primera vez ya que la funcion se maneja en memoria
	return func() int {
		secuencia++
		return numero * secuencia
	}
}
