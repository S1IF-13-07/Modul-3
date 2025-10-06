package main

import "fmt"

func main(){
	var tahun int
	fmt.Printf ("Masukkan tahun: ")
	fmt.Scan(&tahun)

	if (tahun % 4 == 0 && tahun % 100 != 0) || (tahun % 400 == 0){
		fmt.Printf ("kabisat : ")
		fmt.Println(true)
	} else {
		fmt.Printf ("kabisat : ") 
		fmt.Println(false)
	}
}