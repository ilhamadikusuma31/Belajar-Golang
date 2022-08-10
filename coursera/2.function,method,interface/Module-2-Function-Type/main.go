package main

import "fmt"

// nested func
// pertama liat ini => func(float64) float64
// kalo dah dapet hasilnya misal x dari perhitungan pertama
// kedua liat ini genDisplayFn(a, Vo, So float64) x
// jadi ngereturn 1 1
func genDisplayFn(a, Vo, So float64) func(float64) float64 {
	return func(waktu float64) float64 {
		return 0.5*a*waktu*waktu + Vo*waktu + So*waktu
	}
}

func main() {
	fn := genDisplayFn(10, 2, 1)
	fmt.Println(fn(3))
	fmt.Println(fn(5))

}
