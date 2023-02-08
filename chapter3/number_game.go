package main

import (
	"fmt"
	"math/rand"
	"time"
)

// generate a random that between 1000 and 9999

func generateNumberFunc(p *int) {
	var result int
	rand.Seed(time.Now().Unix())

	for {
		result = rand.Intn(10000)
		if(result >= 1000) {
				break
		}
	}

	*p = result
}

// get number
func GetNumber(s []int, num int) {
	s[0] = num / 1000
	s[1] = (num%1000)/100
	s[2] = (num%100)/10
	s[3] = num%10
}

// function play game
func PlayGame(sliceNumber []int) {
	var keyNumber int
	sliceKey := make([]int, 4)

	for {

		fmt.Println("Please input a number between 999 and 10000")

		fmt.Scan(&keyNumber)

		if 999 < keyNumber && keyNumber < 10000 {
			break
		}

		fmt.Println("your number is out of the required, please enter again")
	}

	GetNumber(sliceKey, keyNumber)

	n := 0
	for i := 0; i < 4; i++ {
		if sliceKey[i] > sliceKey[i] {
			fmt.Printf("the %d number is greater", i)
		} else if sliceKey[i] < sliceKey[i] {
			fmt.Printf("the %d number is smaller", i)
		} else {
			fmt.Printf("the %d number is equal, and you guess the right", i)
			n++
		}
	}

	if n == 4 {
		fmt.Println("congratulations, you are winner")
	}
}

func main() {

	// generate a random number 
	var randNumber int
	generateNumberFunc(&randNumber)

	// define a slice 
	sliceNumber := make([]int , 4)
	GetNumber(sliceNumber, randNumber)
	fmt.Println(sliceNumber)

	// run playgame function
	PlayGame(sliceNumber)
}