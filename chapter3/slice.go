package main

import "fmt"

func updateSlice(arr []int) {
	arr[0] = 100
}

func printSlice(arr []int) {
	fmt.Printf("len = %d, cap=%d, arr=%v \n", len(arr), cap(arr), arr)
}

func main() {

	// initial an array
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	// slice
	s := arr[2:6]
	k := arr[:6]
	m := arr[:]
	d := arr[2:]

	fmt.Printf("this is %v %v %v %v \n", s, k, m, d)
	updateSlice(s)
	fmt.Printf("this is %v %v %v %v \n", s, k, m, d)
	fmt.Printf("this is %v \n", arr)

	// append slice
	a := append(s, 11)
	printSlice(a)

	// make to generate slice
	s2 := make([] int, 16)
	s3 := make([] int, 10, 32)
	printSlice(s2)
	printSlice(s3)

	// copy slice
	copy(s3, s)
	printSlice(s3)

	// delete element from slice
	s3 = append(s3[:3], s3[4:]...)
	printSlice(s3)
}