package main

import "fmt"

type student struct {
	id   int
	name string
	gpa  float32
}

func main() {
	var student1 student

	student1.id = 202044055
	student1.name = "vinyl"
	student1.gpa = 3.7

	fmt.Println(student1.id)
	fmt.Println(student1.name)
	fmt.Println(student1.gpa)
}
