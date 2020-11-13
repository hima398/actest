package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
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
	w := bufio.NewWriter(file)
	w.WriteString(strconv.Itoa(n) + "\n")
	for i := 0; i < n; i++ {
		v := rand.Int() % max
		w.WriteString(strconv.Itoa(v) + " ")
	}
	w.Flush()
}
