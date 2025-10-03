package main

import "fmt"

func main(){
	var luas, alas, tinggi float64
	fmt.Print("Masukkan Nilai Alas Dan Tinggi: ")
		fmt.Scan(&alas, &tinggi)

		luas = 0.5 * alas * tinggi
		fmt.Printf("Hasil Luas Segitiga: %.1f", luas)
}