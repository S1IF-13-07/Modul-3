package main

import "fmt"

func main (){
	var luas, alas, tinggi float64

	fmt.Print("Masukan nilai alas segitiga: ")
	fmt.Scan(&alas)
	fmt.Print("Masukan nilai tinggi segitiga: ")
	fmt.Scan(&tinggi)
	luas = 0.5 * alas * tinggi
	fmt.Println("luas segitiga: ", luas)
}