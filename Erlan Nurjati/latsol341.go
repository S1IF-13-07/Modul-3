package main

import "fmt"

func main() {
	var fx float64
	fmt.Print("Masukkan f(x): ")
	fmt.Scan(&fx)
	x := (2 / (fx-5) ) - 5
	fmt.Printf("Hasil persamaannya adalah: %.0f\n", x)
}
