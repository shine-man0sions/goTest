package main // 程序必须有一个main包

import "fmt"
import "os"

func main() {
	fmt.Println("Hello World")
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = ""
	}
	fmt.Println()
}
