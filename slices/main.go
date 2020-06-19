package main

import "fmt"

//Matriz
//var Matriz [5][7]int

func main() {
	//vector
	/*tabla := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, val := range tabla {
		fmt.Println(val)
	}*/

	//slices
	matriz := []int{2, 5, 4}
	fmt.Println(matriz)
	variante2()
}

func variante2() {
	//el make construira un arreglo de enteros, de 5 espacios pero en memoria de 20
	elementos := make([]int, 0, 0)
	//cap muestra la capacidad de elementos de un arreglo
	fmt.Printf("Largo %d, capacidad %d", len(elementos), cap(elementos))
	for i := 0; i < 10; i++ {
		elementos = append(elementos, i)
	}
	fmt.Printf("Largo %d, capacidad %d", len(elementos), cap(elementos))
}
