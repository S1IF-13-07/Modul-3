package main

import (
	"fmt"
	"math"
)

func main() {
	var jarijari float64

	fmt.Print("Masukkan Jari-jari : ")
	fmt.Scanln(&jarijari)

	var volume = 4.0 / 3.0 * math.Pi * jarijari * jarijari * jarijari
	var luas = 4 * math.Pi * jarijari * jarijari
	fmt.Printf("Volume Lingkaran = %.2f dan Luas Lingkaran = %.2f", volume, luas)
}
