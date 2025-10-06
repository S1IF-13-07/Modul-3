package main

import "fmt"

func main() {
	var k int
	var hasil float64 = float64(k)

	fmt.Print("masukkan nilai sisi kubus = ")
	fmt.Scan(&k)

	hasil = float64(k * k * k) + 0.5

	fmt.Println("volume kubus adalah = ", hasil )
}