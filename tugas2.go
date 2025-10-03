package main

import( 
	"fmt"
	"math"
)
func main() {
	var r float64

	fmt.Print("inputkan jari jari :")
	fmt.Scan(&r)

	volume := (4.0/3) * math.Pi * math.Pow(r,3)
	luas := 4 * math.Pi * math.Pow(r,2)

	fmt.Printf("Jari-jari bola dengan %.f memiliki volume %.4f dan luas kulit %.4f\n", r, volume, luas)
}
