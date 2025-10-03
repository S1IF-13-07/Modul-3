package main

import "fmt"

func main() {
	var sisi, volume float64

	fmt.Print("Masukkan panjang sisi kubus: ")
	fmt.Scan(&sisi)

	volume = (sisi * sisi * sisi) + 0.5
	fmt.Println("Volume kubus =", volume)
}
