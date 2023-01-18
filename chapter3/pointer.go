package main

import "fmt"

func swap(p1, p2 *int) {
	*p1, *p2 = *p2, *p1
}

func main() {

	var variable1 int = 10
	fmt.Println("this is for memory variable1 is", variable1)

	var p1 *int
	p1 = &variable1
	fmt.Println("this is for store the pointer address of variable1", p1)

	*p1 = 111
	fmt.Printf("a = %d, *p = %d \n", variable1, *p1)

	//dynamic assign memory
	var p2 *int
	p2 = new(int)
	fmt.Printf("this is *p2 %v \n", *p2)
	*p2 = 333
	fmt.Printf("this is *p2 %v \n", *p2)

	var a1 int = 101
	var a2 int = 301
	swap(&a1, &a2)
	fmt.Printf("exchange two numbers a1 = %d, a2 = %d \n", a1, a2)
}
