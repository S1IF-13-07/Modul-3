package main

import "fmt"

func main() {
	var r float64 
	var phi float64 = 3.1415926535
	fmt.Print("Masukan jari jari bola = ")
	fmt.Scanln(&r)
	var v = 4.0 / 3.0 * phi * ( r * r * r )
	var l = 4 * phi * ( r * r )
	fmt.Printf("Bola dengan jejari %v memiliki volume %v dan luas kulit %v", r, v, l)
}