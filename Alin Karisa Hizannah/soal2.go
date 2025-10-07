package main

import "fmt"

func main() {
	const phi = 3.1415926535
	var r, volume, luas float64

	fmt.Print("Jejari = ")
	fmt.Scan(&r)

	// Rumus
	volume = (4.0 / 3.0) * phi * r * r * r
	luas = 4 * phi * r * r

	// Output
	fmt.Printf("Bola dengan jejari %.0f memiliki volume %.4f dan luas kulit %.4f\n", r, volume, luas)
}
