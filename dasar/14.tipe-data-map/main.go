package main

import "fmt"

func main() {
	//map[tipeKey][tipeValue]
	person := map[string]string{
		"name": "ilham",
		"umur": "20",
	}

	items := map[string]int{
		"kode":  2131,
		"berat": 40,
	}

	items["kode"] = 03243
	items["kedaluarsa"] = 200522

	fmt.Println(person)
	fmt.Println(items)

	var books map[string]string = make(map[string]string)
	books["title"] = "serlok holms"
	books["berat"] = "500"
	books["ups"] = "salah"
	fmt.Println(books)
	delete(books, "ups")
	fmt.Println(books)
}
