package main

import "fmt"

//nb: slice di golang itu sudah memakai prinsip pointer jadi gausah pake tanda & *

func Swap(arr []int, index int) {
	if arr[index] > arr[index+1] {
		arr[index], arr[index+1] = arr[index+1], arr[index]
	}

}

func BubbleSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			Swap(array, j)
		}
	}
	return array
}
func main() {
	array := []int{1, 23, 5, 43, 234, 402, 0, 65, 91, 90}
	fmt.Println(BubbleSort(array))
}
