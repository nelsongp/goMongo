package main

import (
	"fmt"
	"time"
)

func main() {
	canal1 := make(chan time.Duration)
	go bucle(canal1)
	fmt.Println("Llegue hasta aquí")

	//mientras no se le asigne un valor a msg el programa va a esperar
	msg := <-canal1
	fmt.Println(msg)
}

func bucle(canal1 chan time.Duration) {
	inicio := time.Now()
	for i := 0; i < 100000000000; i++ {

	}
	final := time.Now()
	//sub es el que calcula la funcion de inicio de datos
	canal1 <- final.Sub(inicio)
}
