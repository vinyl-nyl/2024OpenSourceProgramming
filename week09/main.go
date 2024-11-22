package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	target := rand.Intn(5) + 1 // print 0, 1, 2
	fmt.Printf("%d\n", target)
}
