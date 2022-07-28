package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var name_answer string
	var addr_answer string
	fmt.Println("Please Input Your Name: ")
	fmt.Scan(&name_answer)
	fmt.Println("Please Input Your Address: ")
	fmt.Scan(&addr_answer)

	p1 := map[string]string{
		"name": name_answer,
		"addr": addr_answer,
	}

	fmt.Println("\nOutput:")
	p1_arr, _ := json.Marshal(p1)
	fmt.Println(string(p1_arr))

}
