package main

import "fmt"

func main (){
	var tahun int
	fmt.Print("Masukan tahun: ")
	fmt.Scan(&tahun)
	var kabisat bool
	kabisat = (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0)
	fmt.Println("Kabisat: ", kabisat)
}