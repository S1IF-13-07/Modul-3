package main

import "fmt"

func main (){
	var sisi, volume float64
	fmt.Print("Masukan Nilai Sisi Kubus: ")
    fmt.Scan(&sisi)
    volume = sisi * sisi * sisi
    fmt.Println("Volume kubus: ", volume)
}