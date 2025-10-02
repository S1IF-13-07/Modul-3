package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung volume bola
func calculateVolume(r float64) float64 {
	return (4.0 / 3.0) * math.Pi * math.Pow(r, 3)
}

// Fungsi untuk menghitung luas kulit bola
func calculateSurfaceArea(r float64) float64 {
	return 4 * math.Pi * math.Pow(r, 2)
}

func main() {
	var radius float64

	// Meminta input jari-jari dari pengguna
	fmt.Print("Masukkan jari-jari bola: ")
	_, err := fmt.Scan(&radius)
	if err != nil {
		fmt.Println("Input tidak valid. Harap masukkan angka.")
		return
	}

	// Validasi jari-jari tidak boleh negatif
	if radius < 0 {
		fmt.Println("Jari-jari tidak boleh negatif.")
		return
	}

	// Menghitung volume dan luas kulit bola
	volume := calculateVolume(radius)
	surfaceArea := calculateSurfaceArea(radius)

	// Menampilkan hasil dengan 4 digit desimal
	fmt.Printf("Bola dengan jari-jari %.2f memiliki Volume %.4f\n dan luas kulit %.4f\n", radius, volume, surfaceArea)

}