package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var n int
	fmt.Scan(&n)
	max := 100
	for i := 0; i < n; i++ {
		v := rand.Int() % max
		fmt.Printf("%d ", v)
	}
}
