package main

import "fmt"

func main() {
	var bulan = [...]string{
		"jan",
		"feb",
		"mar",
		"apr",
		"mei",
		"juni",
		"juli",
		"agt",
		"sep",
		"okt",
		"nov",
		"des",
	}
	var slice1 = bulan[2:9]
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	var slice2 = bulan[10:]
	fmt.Println(slice2)

	var slice3 = "bulan13"
	fmt.Println(slice3)
}
