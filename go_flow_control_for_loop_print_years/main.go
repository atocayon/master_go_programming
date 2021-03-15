package main

import "fmt"

func main() {
	birthdayYear, currentYear := 1996, 2021

	for birthdayYear <= currentYear {
		fmt.Println(birthdayYear)
		birthdayYear++
	}
}
