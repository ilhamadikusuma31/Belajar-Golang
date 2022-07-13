package main

import "fmt"

func main() {
	type NoKtp string

	var no_ktp_ilham NoKtp = "20202020"
	var no_ktp_ragil NoKtp = "20234234"
	var no_ktp_adi NoKtp = "2342344"
	var no_ktp_ayah NoKtp = "46214745"
	var no_ktp_ibu NoKtp = "24653654"
	fmt.Println(no_ktp_ibu)
	fmt.Println(no_ktp_adi)
	fmt.Println(no_ktp_ayah)
	fmt.Println(no_ktp_ragil)
	fmt.Println(no_ktp_ilham)

}
