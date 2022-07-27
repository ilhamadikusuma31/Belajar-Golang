package main

import (
	"fmt"
	"strings"
)

func cek(answer string) string {
	if answer[0] == 'i' || answer[0] == 'I' && answer[len(answer)-1] == 'n' || answer[len(answer)-1] == 'N' {
		if strings.Contains(answer[1:len(answer)-1], "a") {
			return "Found"
		}
	}

	return "Not Found"
}

func main() {
	var answer string
	fmt.Println("please input a string: ")
	fmt.Scan(&answer)
	fmt.Println(cek(answer))
}
