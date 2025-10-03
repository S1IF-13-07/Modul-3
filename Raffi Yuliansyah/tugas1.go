package main

import "fmt"

func main() {
	var fx, x float64
	fmt.Print("Input nilai f(x) = ")
	fmt.Scanln(&fx)
	x = ((2 / (fx - 5)) - 5)
	fmt.Printf("Ouput nilai dari x = %.0f", x)
}