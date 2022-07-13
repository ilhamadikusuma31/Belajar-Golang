package main

import "fmt"

func main() {
	fmt.Println("string")
	fmt.Println("")
	var namaku string
	namaku = "ilhamadi"
	fmt.Println(namaku)
	namaku = "ADI"
	fmt.Println(namaku)
	fmt.Println("")
	fmt.Println("kita bisa melakukan inisialisasi tanpa deklarasi")
	negara := "indonesia"
	fmt.Println(negara)

	var (
		nama_depan    = "ilham"
		nama_belakang = "adikusuma"
	)

	fmt.Println(nama_depan)
	fmt.Println(nama_belakang)
}
