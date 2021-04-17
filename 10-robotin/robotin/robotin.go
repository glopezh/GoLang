package robotin

import (
	"fmt"
	"math/rand"
	"strconv"
)


func Generarobots(x int,mostrar_colisiones bool) []string {
	var names []string
	letters := [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}


	for len(names) < x{
		name := letters[rand.Intn(25)] + strconv.Itoa(rand.Intn(9))
		if (!existe(names,name)){
			names = append(names,name)
		}else{
			if (mostrar_colisiones){
				fmt.Println("Se presento colisión al generar robot",name)
			}
		}
		}
	return names
	}


func existe(names []string, name string) bool {
	for i :=0; i<len(names);i++{
		if(names[i]== name){
			return true
		}
	}
	return false

}