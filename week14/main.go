package main

import (
	"fmt"
)

func main() {
	ages := make(map[string]int)

	var name string
	var age int

	for {
		fmt.Print("What's your name? (exit to 'q' or 'ㅂ') ")
		fmt.Scanln(&name)
		if name == "q" || name == "ㅂ" {
			break
		}
		fmt.Print("Your age? ")
		fmt.Scanln(&age)

		ages[name] = age
	}

	for name, age := range ages {
		fmt.Printf("%s is %d year old.", name, age)
	}
}
