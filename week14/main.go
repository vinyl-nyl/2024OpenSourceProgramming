package main

import "fmt"

func main() {
	var student1 struct {
		id   int
		name string
		gpa  float32
	}

	student1.id = 202044055
	student1.name = "vinyl"
	student1.gpa = 3.7

	fmt.Println(student1.id)

	var student2 struct {
		id   int
		name string
		gpa  float32
	}

	student2.id = 202044066
	student2.name = "nyl"
	student2.gpa = 3.0

	fmt.Println(student2.id)
}
