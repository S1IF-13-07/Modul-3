package main

import (
	"fmt"
)

// isKabisat memeriksa apakah suatu tahun adalah tahun kabisat.
func isKabisat(tahun int) bool {
	// Sebuah tahun adalah kabisat jika habis dibagi 400,
	// ATAU habis dibagi 4 tetapi tidak habis dibagi 100.
	return tahun%400 == 0 || (tahun%4 == 0 && tahun%100 != 0)
}

func main() {
	var tahun int

	// Meminta input tahun dari pengguna
	fmt.Print("Tahun: ")
	_, err := fmt.Scan(&tahun)
	if err != nil {
		fmt.Println("Input tidak valid. Harap masukkan bilangan bulat.")
		return
	}

	// Memanggil fungsi isKabisat untuk memeriksa dan menampilkan hasilnya
	kabisat := isKabisat(tahun)
	fmt.Printf("Kabisat= %t\n", kabisat)
}