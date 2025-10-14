package main

import "fmt"

func main() {
	var a, t float64
	fmt.Print("Masukkan alas dan tinggi: ")
	fmt.Scan(&a, &t)
	fmt.Printf("Luas segitiga: %.2f\n", 0.5*a*t)
}