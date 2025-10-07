package main

import "fmt"

func main () {
	var sisi int
	fmt.Print("Masukkan panjang sisi: ")
	fmt.Scan(&sisi)
	var nilaiSisi = float64(sisi)
	volume := (nilaiSisi * nilaiSisi * nilaiSisi) + 0.5
	fmt.Printf("Volume kubus: %.1f", volume)
}