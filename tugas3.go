package main
import "fmt"
func main(){
	var tahun int 
	fmt.Print("MASUKAN TAHUN: ")
	fmt.Scan(&tahun)
	if ( tahun % 4 == 0 && tahun %100 !=0) || ( tahun %400 ==0){
		fmt.Print ("KABISAT:")
		fmt.Println(true)
	} else {
		fmt.Printf ("KABISAT:")
		fmt.Println(false)
	}

}