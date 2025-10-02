package main

import "fmt"

func main() {
	var tahun int
	fmt.Print("Masukkan tahun : ")
	fmt.Scan(&tahun)
	
	if (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0) {
		fmt.Printf("Kabisat: ")
		fmt.Println(true)
	} else {
		fmt.Printf("Kabisat: ")
		fmt.Println(false)
	}
}