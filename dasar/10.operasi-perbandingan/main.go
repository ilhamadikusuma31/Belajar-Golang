package main

import "fmt"

func main() {
	type NoKtp string

	var angka1 = 20
	var angka2 = 30

	fmt.Println(angka2 != angka1)
	fmt.Println(angka2 == angka1)
	fmt.Println(angka2 >= angka1)
	fmt.Println(angka2 <= angka1)
	fmt.Println(angka2 > angka1)
	fmt.Println(angka2 < angka1)

}
