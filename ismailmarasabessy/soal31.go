package main

import "fmt"

func main() {
	var x float64
	fmt.Print("Masukkan nilai x : ")
	fmt.Scan(&x)

	x = (2 / (x + 5)) + 5

	fmt.Printf("Hasil f(x) = %g", x)
}
