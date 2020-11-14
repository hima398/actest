package main

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	n := 200000
	max := 100
	interval := 1000
	fileName := "in.txt"
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	w.WriteString(strconv.Itoa(n) + "\n")
	s := ""
	for i := 0; i < n; i++ {
		v := rand.Int() % max
		s += strconv.Itoa(v) + " "
		if i%interval == 0 {
			w.WriteString(s)
			w.Flush()
		}
	}
	w.WriteString(s)
	w.Flush()
}
