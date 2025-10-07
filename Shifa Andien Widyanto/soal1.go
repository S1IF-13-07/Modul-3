package main
import "fmt"
func main() {
	var fx, x float64
	fmt.Print("Masukan nilai fx (bilangan real) : ")
	fmt.Scan(&fx)
	x = 2/(fx-5) - 5
	fmt.Printf("%.0f\n", x)
}
