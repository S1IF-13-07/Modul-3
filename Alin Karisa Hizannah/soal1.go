package main

import "fmt"

func main() {
	var fx, x float64

	fmt.Print("Masukkan nilai f(x): ")
	fmt.Scan(&fx)

	// Rumus dari f(x) = 2/(x+5) + 5
	// Maka x = (2 / (f(x) - 5)) - 5
	x = (2 / (fx - 5)) - 5

	fmt.Printf("Nilai x = %.0f\n", x)
}
