package main

import "fmt"

func main() {
	var tahun int
	var kabisat bool

	fmt.Print("Tahun: ")
	fmt.Scan(&tahun)

	if tahun%400 == 0 {
		kabisat = true
		fmt.Print("kabisat = ", kabisat)
	} else if tahun%4 == 0 {
		kabisat = true
		fmt.Print("kabisat = ", kabisat)
	} else if tahun%100 == 0 {
		kabisat = false
		fmt.Print("kabisat = ", kabisat)
	} else {
		kabisat = false
		fmt.Print("kabisat = ", kabisat)
	}
}