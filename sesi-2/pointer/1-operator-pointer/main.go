package main

import "fmt"

/*
	Pointer di golang memiliki dua operator yaitu address operator dan indirect operator.
	Address operator (&) digunakan untuk mendapatkan alamat memori dari suatu variabel.
	Indirect operator (*) digunakan untuk mendapatkan nilai dari alamat memori yang ditunjuk.
*/

func main() {
	name := "Vania"
	age := 20
	isMarried := true

	//TODO: Dari variable yang telah diberikan, kamu dapat mencoba untuk menggunakan address operator dan indirect operator.
	//start_answers
	pointerName := &name
	pointerAge := &age
	pointerIsMarried := &isMarried

	fmt.Println(&pointerName, *pointerName)
	fmt.Println(&pointerAge, *pointerAge)
	fmt.Println(&pointerIsMarried, *pointerIsMarried)
	//end_answer
}
