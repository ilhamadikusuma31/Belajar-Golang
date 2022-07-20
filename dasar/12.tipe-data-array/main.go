package main

import "fmt"

func main() {
	var nama [4]string
	nama[0] = "ilham"
	nama[1] = "adi"
	nama[2] = "kusuma"
	nama[3] = "hamadi"

	fmt.Println(nama[3])

	var values = [4]int{
		10, 32, 12, 43,
	}

	fmt.Println(values)
	fmt.Println(len(values))

}
