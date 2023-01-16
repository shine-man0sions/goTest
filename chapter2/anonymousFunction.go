package main

import "fmt"

func main() {

	// var tow variables
	a := 10
	str := "mike"

	// 1 -  definite an anonymous functions, run later
	f1 := func() {
		fmt.Printf("only deninite and run later %d, %s \n", a, str)
	}

	// run the anonymous functions
	f1()

	// using type
	type FuncType func()
	var f2 FuncType
	f2 = f1
	f2()

	// 2- definite an anonymous functions and run at the same time

	func() {
		fmt.Printf("definite and run %d, %s \n", a, str)
	}()

	// 3- definite an anonymous functions with params and run later

	f3 := func(a int, b int) {
		fmt.Printf("the %d,  + %d = %d \n", a, b, a+b)
	}

	f3(10, 20)

	// 4 - definite an anonymous functions with params and run at the same time

	func(a, b int) {
		fmt.Printf("the %d,  + %d = %d \n", a, b, a+b)
	}(11, 23)

	// 5 - definite an anonymous functions with params have a return and run at the same time
	x1, x2 := func(a, b int) (min, max int) {
		if (a - b) < 0 {
			min = a
			max = b
		} else {
			min = b
			max = a
		}
		return
	}(11, 23)

	fmt.Printf("the %d is smaller than %d \n", x1, x2)
}
