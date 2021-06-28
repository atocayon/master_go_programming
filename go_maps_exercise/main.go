package main

import "fmt"

func main() {
	// Exercise #1
	var m1 map[string]int
	fmt.Printf("Type of m1: %T\n", m1)

	m2 := map[int]string{1: "One", 2: "Two"}
	m2[3] = "Abba"

	fmt.Println(m2[2])   //existing key
	fmt.Println(m2[100]) //non-existing key

	// Exercise #2

	var m11 map[int]bool
	// m1[5] = true //Error cannot assign a value to uninitialized map
	_ = m11
	m22 := map[int]int{3: 10, 4: 40}
	m33 := map[int]int{3: 10, 4: 40}

	// fmt.Println(m2 == m3) //Error map can only be compared to nil
	_ = m22
	_ = m33

	// Exercise #3
	m := map[int]bool{1: true, 2: false, 3: false}

	delete(m, 2) // Delete a key valu pair

	// Iterate over the map and print out each value
	for k, v := range m {
		fmt.Printf("key: %d value: %v\n", k, v)
	}
}
