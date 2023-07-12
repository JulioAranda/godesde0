package main

import (
	"fmt"
	"runtime"

	"github.com/JulioAranda/godesde0/variables"
)

func main() {
	//variables.MuestraEnteros()
	//variables.RestoVariables()
	estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)

	if os := runtime.GOOS; os == "linux" || os == "darwin" {
		fmt.Println("Esto es Mac")

	} else {
		fmt.Println("Esto es Windows ")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")

	case "Windows":
		fmt.Println("Esto es Windows ")

	default:
		fmt.Printf("%s \n", os)
	}

}
