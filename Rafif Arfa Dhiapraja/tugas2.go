package main

import "fmt"

func main() {
	var volumebola float64
	var luaskulit float64
	var r int
	var pi float64 = 3.1415926535 

	fmt.Print("Masukkan jari-jari bola: ")
	fmt.Scan(&r)


	volumebola = (float64(4) / float64(3)) * pi * float64(r*r*r)
	luaskulit = 4 * pi * float64(r*r)

	fmt.Printf("Bola dengan jejari %d memiliki volume %.4f dan luas kulit %.4f", r, volumebola, luaskulit)
}