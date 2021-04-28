package main

import (
	"11-robotin-espacial/funciones-espaciales"
	"fmt"
)

func main() {
	var issnow = funciones_espaciales.Issnow()
	fmt.Printf("Latitud",issnow[0])
	fmt.Printf("Longitud",issnow[1])
	go funciones_espaciales.Whereisiss()
	//funciones_espaciales.Apigoogle()
}
