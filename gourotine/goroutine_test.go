package main

import (
	"fmt"
	"testing"
	"time"
)

func say() {
	fmt.Println("hai")
}

func displayNumber(number int) {
	fmt.Println("angkaku: ", number)
}

func TestCreateGoroutine(t *testing.T) {
	go say()
	fmt.Println("haiii aja")
}

func TestBanyakGoroutine(t *testing.T) {
	for i := 0; i <= 100000; i++ {
		go displayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
