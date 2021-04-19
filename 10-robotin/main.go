package main

import (
	"fmt"
	"10-robotin/robotin"
	"time"
)

func main() {
	var nulo [] string
	var nombres = robotin.Generarobots(100,false,nulo)
	fmt.Println("Mis 100 robots")
	fmt.Println(nombres)
	for i:=1; i>=0;i++ {
		fmt.Println("Grupo",i)
		var nuevos_nombres = robotin.Generarobots(25,true,nombres)
		fmt.Println(nuevos_nombres)
		time.Sleep(1 * time.Second)
	}
}

