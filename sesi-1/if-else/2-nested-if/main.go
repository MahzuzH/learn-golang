package main

import "fmt"

/*
	Disini kita akan belajar nested-if, yaitu if didalam if

	Studi kasus:
		Buatlah program sederhana
		Jika jenis kelamin P maka cetak perempuan, jika jenis kelamin L maka cetak laki-laki
		Jika nilai >= 80 cetak lulus, jika nilai < 80 cetak tidak lulus

	Contoh Input :
		jenis_kelamin = L
		nilai = 85

	Contoh Output :
		laki-laki dan lulus
*/

func main() {
	jenis_kelamin := "L"
	nilai := 90

	//T
	//start_answer
	if jenis_kelamin == "L" {
		if nilai >= 80 {
			fmt.Println("laki laki dan lulus")
		} else {
			fmt.Println("laki laki dan tidak lulus")
		}
	} else if jenis_kelamin == "P" {
		if nilai >= 80 {
			fmt.Println("perempuan dan lulus")
		} else {
			fmt.Println("perempuan dan tidak lulus")
		}
	}
	//end_answer
}
