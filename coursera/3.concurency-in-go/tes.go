package main

import "fmt"

func perkalian(v1, v2 int, ch chan int) {
	ch <- v1 * v2
}

func main() {
	c := make(chan int, 9)
	go perkalian(1, 2, c)
	go perkalian(3, 4, c)
	a := <-c
	b := <-c
	fmt.Println(a * b)
}
