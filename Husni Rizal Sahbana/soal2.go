package main

import (
	"fmt"
	"math"
)

func main() {
	var jari_jari int
	fmt.Print("Masukkan jari-jari bola : ")
	fmt.Scan(&jari_jari)

	// konversi ke float64 agar bisa digunakan di perhitungan
	radius := float64(jari_jari)

	// hitung volume dan luas
	volume := (4.0 / 3.0) * math.Pi * math.Pow(radius, 3)
	luas := 4.0 * math.Pi * math.Pow(radius, 2)

	// tampilkan hasil
	fmt.Printf("Volume bola: %.4f\n", volume)
	fmt.Printf("Luas permukaan bola: %.4f\n", luas)
}
