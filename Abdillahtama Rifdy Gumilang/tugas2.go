package main

import "fmt"

func main() {
	const phi = 3.1415926535
	var jejari float64
	fmt.Print("Masukkan Jejari = ")
	fmt.Scan(&jejari)
	volume := (4.0 / 3.0) * phi * jejari * jejari * jejari
	luas := 4 * phi * jejari * jejari
	fmt.Printf("Bola dengan jejari %.0f memiliki volume %.4f dan luas kulit %.4f", jejari, volume, luas)
}
