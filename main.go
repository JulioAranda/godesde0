package main

import (
	"fmt"

	"github.com/JulioAranda/godesde0/variables"
)

func main() {
	//variables.MuestraEnteros()
	//variables.RestoVariables()
	estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
}
