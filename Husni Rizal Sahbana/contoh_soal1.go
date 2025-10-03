package main

import "fmt"

func main() {

	var sisi_kubus int
	fmt.Printf("Masukan panjang sisi kubus : ")
	fmt.Scan(&sisi_kubus)
	volume := sisi_kubus * sisi_kubus * sisi_kubus
	fmt.Print("Volume sisi kubus adalah : ", volume)
}
