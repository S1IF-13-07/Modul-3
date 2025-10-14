package main

import "fmt"

func main() {
	var s int
	fmt.Print("Sisi: ")
	fmt.Scan(&s)
	fmt.Println("Volume:", s*s*s)
}