package main

import "fmt"

// Sekarang kita akan coba untuk membaca huruf pertama dan terakhir dari nama kalian
// Buatlah nama kalian menjadi arrat, contoh: [R, O, B, E, R, T]
// Dengan mengakses indeks 0 untuk huruf pertama, dan indeks 5 untuk huruf terakhir
// Didaptkan hasil R dan T

func main() {
	//TODO: Buatlah nama kalian menjadi array
	//start_answer
	nama := [6]string{"M","A","H","Z","U","Z"}
	//end_answer

	//TODO: Cetak huruf pertama dari nama kalian dengan mengakses indeks 0
	//start_answer
	fmt.Println(nama[0])
	//end_answer

	//TODO: Cetak huruf terakhir dari nama kalian dengan mengakses indeks terakhir
	//start_answer
	fmt.Println(nama[5])
	//end_answer

	//TODO: dengan menggunakan perintah loop cetak nama kalian secara berurutan
	//Hint: gunakan for range
	//Hint: gunakan fmt.printf() untuk mencetak tambah menambahkan baris baru
	//start_answer
	for i:=0 ; i<= len(nama); i++ {
		fmt.Println(nama[i])
	}
	//end_answer
}
