package main

import (
	"fmt"
	"math"

)

func main (){
	var r, volume, luas float64
	fmt.Print("Masukan jari-jari bola: ")
	fmt.Scan(&r)
	volume = 4 * math.Pi * r * r * r / 3
	luas = 4 * math.Pi * r * r 
	fmt.Printf("Bola dengan jari-jari %.0f memiliki volume %.4f dan luas kulit %.4f\n", r, volume, luas) 
}