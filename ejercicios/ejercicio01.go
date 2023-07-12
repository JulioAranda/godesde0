package ejercicios

import (
	"fmt"
	"strconv"
)

func EvaluaEntero(valor string) (int, string) {
	entero, _ := strconv.Atoi(valor)

	/*
		entero, error := strconv.Atoi(valor)
		if error != nil {
			fmt.Println("Error de conversion")
		}
	*/

	if entero > 100 {
		fmt.Printf("El entero[%d] es mayor a 100 \n", entero)

	} else {
		fmt.Printf("El entero[%d] es menor igual 100 \n", entero)
	}

	return entero, valor
}
