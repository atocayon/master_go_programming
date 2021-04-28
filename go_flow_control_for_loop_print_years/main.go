package main

import "fmt"

func main() {
	birthdayYear, currentYear := 1996, 2021

	for i := birthdayYear; i <= currentYear; i++ {
		fmt.Printf("%d ", i)
	}
}
