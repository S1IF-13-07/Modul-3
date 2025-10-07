package main

import "fmt"

func main() {
	//INPUT SISI
	var sisi int
	fmt.Print("Masukkan sisi kubus: ")
	fmt.Scan(&sisi)
	
	//HITUNG VOLUME
	var volume float64
	volume = (float64(sisi))*(float64(sisi))*(float64(sisi)) + 0.5
	

	fmt.Printf("Volume Kubus = %.1f\n", volume)
}
