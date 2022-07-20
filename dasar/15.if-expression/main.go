package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan Nama: ")
	text, _ := reader.ReadString('\n')

	if text == "ilham" {
		fmt.Println("bener broh")
	} else if text == "adi" {
		fmt.Println("haiiiiii")
	} else {
		fmt.Println("salah!!")
	}
}
