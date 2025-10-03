package main

import "fmt"

func main() {
	var sisi int
	fmt.Print("Masukkan panjang sisi kubus : ")
	fmt.Scan(&sisi)
	hasil := sisi * sisi * sisi
	volume := float64(hasil) + 0.5
	fmt.Printf("Volume dari adalah kubus adalah %.1f", volume)
}
