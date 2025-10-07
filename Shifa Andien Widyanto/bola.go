package main

import "fmt"
func main(){
	var r float64

	fmt.Print("Jejari = ")
	fmt.Scan(&r)

	var v = (4.0/3.0) * 3.1415926535 * r * r * r
	var l = 4.0 * 3.1415926535 * r * r

	fmt.Printf("Bola dengan jejari %d, memiliki volume, %.4f, dan luas kulit, %.4f", int(r), v, l)
}