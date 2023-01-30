package main

import "fmt"

/**
syntax make(map[string] int)
get element m[key]
value, ok := m[key]
*/

func main() {

	// define a map
	m := map[string]string {
		"name" : "YiHang",
		"course" : "ccs",
		"site" : "github",
		"quality" : "high", 
	}

	fmt.Printf("map = %v \n", m)

	// define an empty map
	// m2 == empty map
	m2 := make(map[string] int)

	// m3 == nil
	var m3 map[string] int
	fmt.Println(m2)
	fmt.Println(m3)


	// key do not have sort, element will out put random
	for k, v := range m {
		fmt.Printf("key = %s, value = %s\n", k, v)
	} 

	// get value using map
	if name, ok := m["name"]; ok {
		fmt.Println(name)
	} else {
		fmt.Println("key does not exist")
	}

	name, ok := m["name"]
	fmt.Println(name, ok)
	// delete values
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}