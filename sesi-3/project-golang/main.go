package main

import (
	"fmt"
	"project-golang/student"
)

func main() {
	budi := student.Student{
		NIM:  "123",
		Name: "Budi",
		Age:  20,
		IPK:  4.1,
	}

	status, err := budi.GraduationStatus()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(status)
}
