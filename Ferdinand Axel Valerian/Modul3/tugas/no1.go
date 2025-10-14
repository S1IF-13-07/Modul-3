package main

import (
	"fmt"
	"math"
)

func main() {
	var fx float64
	fmt.Print("Masukkan nilai f(x): ")
	fmt.Scan(&fx)
	hasil := (2 / (fx - 5)) - 5
	x := int(math.Ceil(hasil))
	fmt.Printf("Nilai x adalah: %d", x)
}
