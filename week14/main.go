package main

import "fmt"

type visitor struct {
	age  int
	cost int
}

func calculateCost(visitors []visitor) int {
	totalCost := 0
	for _, v := range visitors {
		totalCost = totalCost + v.cost
	}
	return totalCost
}

func main() {
	var numVisitors int
	fmt.Println("How many visitors : ")
	fmt.Scanln(&numVisitors)

	var vs []visitor
	vs = make([]visitor, numVisitors)

	for i := 0; i < numVisitors; i++ {
		var age int
		fmt.Print("Input age : ")
		fmt.Scan(&age)

		switch {
		case age < 12:
			vs[i] = visitor{age: age, cost: 5}
		case age >= 12 && age < 65:
			vs[i] = visitor{age: age, cost: 10}
		case age >= 65:
			vs[i] = visitor{age: age, cost: 7}
		}
	}

	fmt.Printf("Total price is %d$.", calculateCost(vs))
}
