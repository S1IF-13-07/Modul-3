package main

import "fmt"

func main() {
	var fx float64
	fmt.Print("Masukkan Nilai : ")
	fmt.Scanln(&fx)

	var x = (2 / (fx - 5)) - 5
	fmt.Printf("Hasilnya adalah %.3f\n", x)
}
