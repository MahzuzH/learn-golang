package main

import "fmt"

// Assigment kali ini dinamakan fizzbuzz
// Tuliskan program for loop dari 1 sampai dengan 100.
// lalu setiap perulangan program ini.

// Jika angka tersebut habis dibagi 3, maka tampilkan "fizz"
// Jika angka tersebut habis dibagi 5, maka tampilkan "buzz"
// Jika angka tersebut habis dibagi 3 dan 5, maka tampilkan "fizzbuzz"

// Tips gunakan % (modulo) untuk mengetahui apakah angka tersebut habis dibagi 3 atau 5 atau tidak.

func main() {
	for i := 1; i <= 100; i++ {
		//TODO
		//start_answer
		if i % 3 == 0 && i % 5 == 0{
			fmt.Println(i ,"fizzbuzz")
		} else if i % 3 == 0{
			fmt.Println(i, "fizz")
		} else if i % 5 == 0{
			fmt.Println(i, "buzz")
		}
		//end_answer
	}
}
