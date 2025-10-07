package main
 import "fmt"

 func main() {
	var sisi int
	var volume float64
	fmt.Print("Masukkan panjang sisi kubus: ")
	fmt.Scan(&sisi)
	volume = (float64(sisi)) * (float64(sisi)) * (float64(sisi))
	fmt.Println("Volume kubus dengan sisi =", volume)
 }