package main

import "fmt"

func main() {
	var answer float32
	fmt.Println("Converter Floating Point to Int\n please insert a floating point: ")
	fmt.Scan(&answer)
	fmt.Println(int(answer))
}
