package main

import "fmt"

func main() {
	var sisi, volume float64
	fmt.Print("masukkan panjang sisi kubus: ")
	fmt.Scan(&sisi)
	volume = sisi*sisi*sisi + 0.5 //kemaren suruh ditambahin 0.5 kak
	fmt.Print("volume kubus adalah: ", volume)
}