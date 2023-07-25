package main

import (
	"fmt"
)

// Fungsi juga dapat mengembalikan lebih dari satu value
// TODO: Buatlah fungsi Square
// Parameter: dua buah bilangan
// Output: pangkat dua dari tiap bilangan
// Contoh:
/*
	Square(2,3) --> 4,9
*/
//start_answer
func Square (a,b int) (int, int) {
	resulta := a*a
	resultb := b*b
	return resulta, resultb
}

//end_answer

func main() {
	result1, result2 := Square(3, 2)
	fmt.Printf("%d dan %d\n", result1, result2)
	fmt.Println(Square(9, 8))
}
