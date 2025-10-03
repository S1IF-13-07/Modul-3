package main

import "fmt"

func main() {
	const pi = 3.1415926535
	var r, volumeBola, luasPermukaanBola float64
	fmt.Print("Input nilai r (jari-jari bola) = ")
	fmt.Scanln(&r)
	volumeBola = ((4 * pi * r * r * r) / 3)
	luasPermukaanBola = 4 * pi * r * r
	fmt.Printf("Output volume bola dengan r = %.1f adalah %.4f \n", r, volumeBola)
	fmt.Printf("Output luas permukaan bola dengan r = %.1f adalah %.4f", r, luasPermukaanBola)
}