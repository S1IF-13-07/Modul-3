package main

import "fmt"

func main() {
	var fx float64

	// Input nilai f(x)
	fmt.Print("Masukkan nilai f(x): ")
	fmt.Scan(&fx)

	// Rumus: x = 2 / (f(x) - 5) - 5
	x := (2 / (fx - 5)) - 5

	// Tampilkan hasil
	fmt.Println("Nilai x adalah: ", x)
}
