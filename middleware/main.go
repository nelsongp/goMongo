package main

import "fmt"

/*Son interceptores que permiten ejecutar instrucciones comunes a varias funciones que reciben y devuelven los mismos tipos de variables*/

var result int

func main() {
	fmt.Println("inicio")
	result = operacionesMidd(sumar)(2, 3)
	fmt.Println(result)
	result = operacionesMidd(restar)(2, 3)
	fmt.Println(result)
	result = operacionesMidd(multiplicar)(2, 3)
	fmt.Println(result)
}

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

//solo se pone el tipo no es necesario ponerle el valor ejemplo a int, basta solo con poner int
//lo que recibe el middleware debe ser lo mismo que lo que retoran
func operacionesMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Incio de operacion")
		return f(a, b)
	}
}
