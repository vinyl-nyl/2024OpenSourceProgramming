package main

import (
	"fmt"
)

func main() {
	// i := 13
	// var f float64 = 12.9
	// c1 := 'A' // ASCII code
	// c2 := '김' // UNICODE

	// fmt.Printf("value i : %d, value f : %f\n", i, f)
	// // fmt.Printf("value i : %d, value f : %f\n", i, f) // invalid operation: i * f (mismatched types int and float
	// fmt.Printf("%d * %f = %d\n", i, f, i*int(f)) // conversion
	// // fmt.Printf("%d * %f = %f\n", i, f, float64(i)*f) // conversion
	// fmt.Println(c1, c2)
	// fmt.Println(reflect.TypeOf(i), reflect.TypeOf(f), reflect.TypeOf(c1), reflect.TypeOf(c2))

	var f float64
	var i int
	var b bool
	var s string

	fmt.Println(f, i, b, s)
	fmt.Printf("%f%t%s%d\n", f, b, s, i) // zero value

	f = 2.7
	i = 3
	fmt.Print("\n\n", f > float64(i), "\n") // comparsion (true/false)
}
