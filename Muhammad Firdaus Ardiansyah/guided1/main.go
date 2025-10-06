package main

import "fmt"

func main(){
	var sisi int
	var volume float64
	fmt.Printf ("Masukkan panjang sisi kubus: ")
	fmt.Scan(&sisi)
	volume = float64 (sisi * sisi * sisi) + 0.5
	fmt.Printf ("Volume kubus adalah %.1f ", volume)

} 