 package main
import "fmt"
 func main() {
	var celcius float64



	fmt.Print("TEMPERATUR CELCIUS:")
	fmt.Scan(&celcius)
	 

	reamur := (4.0/5.0) * celcius
	fahrenheit := (9.0/5.0) * celcius + 32
	kelvin := celcius + 273


	fmt.Printf("DERAJAT REAMUR: %.f\n", reamur)
	fmt.Printf("DERAJAT FAHRENHEIT: %.f\n", fahrenheit)
	fmt.Printf("DERAJAT KELVIN: %.f\n", kelvin)

}