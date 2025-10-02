package main

import (
	"errors"
	"fmt"
	"math"
)

// Fungsi untuk menghitung f(x)=(2/(x+5))+5 dan membulatkan hasilnya ke bilangan bulat
func calculateAndRound(x float64) (int, error) {
	// Validasi untuk menghindari pembagian dengan nol
	if x+5 == 0 {
		return 0, errors.New("error: pembagian dengan nol karena x + 5 = 0")
	}

	// Lakukan perhitungan
	result := (2 / (x + 5)) + 5

	// Bulatkan hasil ke bilangan bulat terdekat dan ubah ke tipe int
	roundedResult := int(math.Round(result))

	return roundedResult, nil
}

func main() {
	var x float64

	fmt.Print("Masukkan nilai x: ")
	_, err := fmt.Scan(&x)
	if err != nil {
		fmt.Println("Input tidak valid. Harap masukkan angka.")
		return
	}

	roundedResult, err := calculateAndRound(x)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Hasil keluaran f(x)=(2/(x+5))+5 adalah: %d\n", roundedResult)
	}
}