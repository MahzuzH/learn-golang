package main

import "fmt"

/*
	Disini kita akan belajar operator logic
	Ada tiga jenis logical operator dalam Go

		&& (and)
		|| (or)
		! (not)

	Studi kasus:
		Buatlah program sederhana
		Jika jenis kelamin P maka cetak perempuan, jika jenis kelamin L maka cetak laki-laki, jika tidak keduanya maka cetak tidak valid
		Jika nilai diantara 85-100 maka cetak = A
		Jika nilai diantara 70-84 maka cetak = B
		Jika nilai diantara 0 - 70 maka cetak = C

	Contoh Input :
		- Case #1
			jenis_kelamin = L
			nilai = 85
		- Case #2
			jenis_kelamin = X
			nilai = 90

	Contoh Output :
		- Case #1
			laki-laki dan nilai A
		- Case #2
			tidak valid
*/

func main() {
	jenis_kelamin := "P"
	nilai := 90

	//
	//start_answer
	if jenis_kelamin == "L" || jenis_kelamin == "l" {
		if nilai >= 85 && nilai <= 100 {
			fmt.Println("laki laki dan nilai A")
		} else if nilai >= 70 && nilai < 85{
			fmt.Println("laki laki dan nilai B")
		} else if nilai < 70 {
			fmt.Println("laki laki dan nilai C")
		}
	} else if jenis_kelamin == "P" || jenis_kelamin == "p" {
		if nilai >= 85 && nilai <= 100 {
			fmt.Println("perempuan dan nilai A")
		} else if nilai >= 70 && nilai < 85{
			fmt.Println("perempuan dan nilai B")
		} else if nilai < 70 {
			fmt.Println("perempuan dan nilai C")
		}
	} else {
		fmt.Println("tidak valid")
	}
	//end_answer
}