package main

import "fmt"

func main() {
	const π = 3.1415926535
	var r float64

	fmt.Print("Masukkan jejari bola: ")
	fmt.Scan(&r)

	fmt.Printf(
		"Jejari %.0f → Volume: %.4f, Luas: %.4f\n",
		r, (4.0/3.0)*π*r*r*r, 4*π*r*r,
	)
}