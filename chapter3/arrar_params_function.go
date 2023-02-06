package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 数组作为函数参数传递，是值传递
// 数组越大，运行越慢

// 生成随机数组函数
func generateRandomFunc() [10]int {
	rand.Seed(time.Now().UnixNano())
	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = rand.Intn(300)
	}
	return arr
}

func calculateArr(arr [10]int) [10]int {
	for i := 0; i < 10; i++ {
		arr[i] = arr[i] * arr[i]
	}
	return arr
}

func calculateArrPointer(arr *[10]int) {
	for i := 0; i < 10; i++ {
		(*arr)[i] = arr[i] * arr[i]
	}
}

func main() {
	// 生成一个随机数组
	b := generateRandomFunc()
	d := generateRandomFunc()
	c := calculateArr(b)
	fmt.Printf("this is %v \n", b)
	fmt.Printf("this is %v \n", c)

	calculateArrPointer(&d)

	fmt.Printf("this is %v \n", d)
	fmt.Printf("this is %v \n", d)
}
