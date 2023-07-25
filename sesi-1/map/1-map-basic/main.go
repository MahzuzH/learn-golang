package main

import "fmt"

// Disini kita akan belajar tentang map
// map itu kalian bisa anggap sebagai kamus, terdapat kata:definisi
// atau dapat dikatakan key:value

func main() {
	// Contoh inisalisasi map
	/*
		map[type]type
		var m0 map[string]int{}
		m1 := make(map[string]int)
		m := map[string]int{"abc": 123, "xyz": 789} <- langsung di isi dengan key dan value
	*/

	//TODO: buatlah sebuah map kamus kata, yang berisi kata_indonesia:kata_inggirs, buatlah sebanyak 5 kata
	//Map:
	/*
		"mobil":    "car",
		"sekolah":  "school",
		"bangunan": "building", 
		"sakit":    "sick",
		"senang":   "happy",
	*/
	//start_answer
	kamus := map[string]string{
		"mobil":"car", 
		"sekolah":"school", 
		"bangunan":"building", 
		"sakit":"sick", 
		"senang":"happy",
	}
	//end_answer

	//Untuk mengakses value map kita butuh tau key dari value pada map
	//TODO: cetak value dari key "mobil"
	//start_answer
	fmt.Println(kamus["mobil"])
	//begin_answer

	//Kita juga bisa mengecek suatu key pada map ada atau tidak
	//TODO: cek apakah value dari "sedih" ada pada map, jika tidak cetak "kata tidak terdapat pada kamus"
	//start_answer
	value, isExist := kamus["sedih"]
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("kata tidak terdapat pada kamus")
	}

	//Kalian masih ingan 'for range', kita akan mencetak seluruh isi map dengan for-rang
	//TODO cetak seluruh nilai dari map
	//Output:
	/*
		mobil:car
		sekolah:school
		bangunan:building
		sakit:sick
		senang:happys
	*/
	//start_answer
	for key, value := range kamus {
		fmt.Println(key, value)
	}
	//end_answer
}
