package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64

	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scan(&r)

	volume := (4.0 / 3.0) * math.Pi * math.Pow(r, 3)
	luas := 4 * math.Pi * math.Pow(r, 2)

	fmt.Printf("Bola dengan jari-jari %.f memiliki volume %.4f dan luas kulit %.4f\n", r, volume, luas)
}