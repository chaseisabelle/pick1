package main

import (
	"math/rand"
	"os"
	"time"
)

func main() {
	input := os.Args[1:]
	count := len(input)

	if count < 1 {
		panic("i need at least 1 to pick")
	}

	rand.Seed(time.Now().Unix())

	println(input[rand.Intn(count)])
}