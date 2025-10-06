package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64

	fmt.Print("masukkan jari-jari bola = ")
	fmt.Scan(&r)

	vol := (4.0/3.0) * math.Pi * math.Pow(r, 3)
	luas := 4 * math.Pi * math.Pow(r, 2)

	fmt.Printf("Bola dengan jari-jari %.f, memiliki volume %.4f dan luas kulit %.4f\n ", r, vol, luas)
}