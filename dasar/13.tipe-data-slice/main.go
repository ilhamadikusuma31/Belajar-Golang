package main

import "fmt"

func main() {
	var bulan = [...]string{
		"jan", "feb", "mar", "apr", "mei", "juni", "juli", "agt", "sep", "okt", "nov", "des",
	}
	var slice1 = bulan[2:9]
	//fmt.Println(len(slice1))
	//fmt.Println(cap(slice1))

	var slice2 = bulan[:4]
	//fmt.Println(slice2)

	var slice3 = append(slice2, "bulan13")
	//fmt.Println(slice3)

	slice3[1] = "notDesember"
	fmt.Println("ori    : ", bulan)
	fmt.Println("slice 1: ", slice1)
	fmt.Println("slice 2: ", slice2)
	fmt.Println("slice 3: ", slice3)

	fmt.Println("\nMake Slice")
	//                    arr,len,cap
	sliceBaru := make([]string, 2, 5)
	fmt.Println("sliceBaru: ", sliceBaru)
	sliceBaru[0] = "e1"
	sliceBaru[1] = "e2"
	fmt.Println("sliceBaru: ", sliceBaru)

	fmt.Println("\nMake Copy")
	sliceCopyan := make([]string, len(sliceBaru), cap(sliceBaru))
	copy(sliceCopyan, sliceBaru)
	fmt.Println("slice copy: ", sliceCopyan)

	fmt.Println("\nperbedaan slice dan array")
	iniArray1 := [...]int{1, 2, 23, 45, 23, 43}
	iniArray2 := [6]int{1, 2, 23, 45, 23, 43}
	iniSlice := []int{1, 2, 23, 45, 23, 43}
	fmt.Println(iniArray1)
	fmt.Println(iniArray2)
	fmt.Println(iniSlice)

}
