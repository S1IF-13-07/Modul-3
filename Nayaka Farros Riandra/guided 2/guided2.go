package main

import "fmt"

func main() {
	var alas, tinggi, luas float64

	fmt.Print("masukkan nilai alas = ")
	fmt.Scan(&alas)

	fmt.Print("masukkan nilai tinggi = ")
	fmt.Scan(&tinggi)

	luas = 0.5 * alas * tinggi

	fmt.Println("luas balok tersebut adalah = ", luas)
}