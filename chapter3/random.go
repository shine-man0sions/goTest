package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// get the current time
	rand.Seed(time.Now().UnixNano())

	// initial an array
	var arr [10]int

	// generate random
	for i := 0; i < 10; i++ {

		// generate random that smaller than 100
		arr[i] = rand.Intn(100)
	}

	fmt.Printf("arr = %v \n", arr)
}
