package main

import (
	"fmt"
	"time"

	us "./user"
)

//ejemplo de herencia
type pepe struct {
	us.Usuario
}

func main() {
	u := new(pepe)
	u.AltaUsuario(1, "Pablo", time.Now(), true)
	fmt.Println(u.Usuario)
}
