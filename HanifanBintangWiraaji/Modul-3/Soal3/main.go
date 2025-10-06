package main 

import "fmt"

func main() {
	var kabisat bool = false
	var tahun int
	fmt.Print("Tahun = ")
	fmt.Scanln(&tahun)
	if ( tahun % 4 == 0 && tahun % 100 != 0 ) || ( tahun % 400 == 0 ) {
		kabisat = !kabisat
	}
	fmt.Printf("Kabisat = %v", kabisat)
}