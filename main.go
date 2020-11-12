package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	var n int
	fmt.Scan(&n)
	max := 100
	file, err := os.OpenFile("./in.txt", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return
	}
	defer file.Close()
	for i := 0; i < n; i++ {
		v := rand.Int() % max
		fmt.Printf("%d ", v)
	}
}
