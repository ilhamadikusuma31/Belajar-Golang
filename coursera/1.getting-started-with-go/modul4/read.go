package main

import (
	"bufio"
	"fmt"
	"os"
)

type Name struct {
	fname string
	lname string
}

func main() {

	//dat, err := ioutil.ReadFile("test.txt")
	//if err != nil {
	//	fmt.Println("errors:", err)
	//}
	//fmt.Println(string(dat))

	var answer_fname string
	var answer_lname string
	var answer_filename string
	fmt.Println("Please input your firstname: ")
	fmt.Scan(&answer_fname)
	fmt.Println("Please input your lastname: ")
	fmt.Scan(&answer_lname)
	fmt.Println("Please input file name: ")
	fmt.Scan(&answer_filename)

	file, err := os.Open(answer_filename)
	if err != nil {
		fmt.Println("\nFile could not be opened. Please make sure you have the")
		fmt.Println("correct filename and path before running again.")
	}

	container := make([]Name, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		obj := Name{
			fname: answer_fname,
			lname: answer_lname,
		}
		container = append(container, obj)
	}

	fmt.Println("\nHere are your list of names from", answer_filename)
	for i := 0; i < len(container); i++ {
		fmt.Println(container[i].fname, container[i].lname)
	}
	fmt.Println(err)
}
