package main

import "fmt"

func main() {
	var tahun int
	var kabisat bool
	fmt.Print("Input, cek tahun kabisat : ")
	fmt.Scanln(&tahun)
	kabisat = (tahun%400 == 0) || ((tahun%100 != 0) && (tahun%4 == 0))
	fmt.Print("Output, tahun kabisat : ", kabisat)
}