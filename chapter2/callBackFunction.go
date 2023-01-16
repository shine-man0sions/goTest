package main

import "fmt"

// define a function type called FuncType

type FuncType func(int, int) int

// algorithm add function
func addFunc(a, b int) int {
	return a + b
}

// algorithm minus function
func minusFunc(a, b int) int {
	return a - b
}

// algorithm multiply function
func multiplyFunc(a, b int) int {
	return a * b
}

// algorithm divide function
func divideFunc(a, b int) int {
	return a / b
}

// the calculate function

func CalculateFunc(a, b int, callBackFunc FuncType) (result int) {
	result = callBackFunc(a, b)
	return
}

func main() {
	a := 10
	b := 20
	c := CalculateFunc(a, b, addFunc)
	fmt.Printf("a = %d, b = %d , a + b = %d\n", a, b, c)
}
