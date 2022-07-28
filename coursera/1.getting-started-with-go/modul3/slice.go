package main

import "fmt"

func main() {

	var answer int
	foo := make([]int, 3)
	for i := 0; i < len(foo); i++ {
		fmt.Scan(&answer)
		foo[i] = answer
	}
	for i := 0; i < len(foo); i++ {
		temp := foo[i]
		if i == len(foo)-1 {
			continue
		}
		if foo[i] > foo[i+1] {
			foo[i] = foo[i+1]
			foo[i+1] = temp
		}
	}
	fmt.Println("\nsorted:")
	for i := 0; i < len(foo); i++ {
		fmt.Println(foo[i])
	}

}
