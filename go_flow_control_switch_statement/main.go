package main

import "fmt"

func main() {

	// Change the code and use a switch statement instead of an if statement.
	switch age := -9; true {
	case age < 0 || age > 100:
		fmt.Println("Invalid Age")
		break
	case age < 18:
		fmt.Println("You are minor!")
		break
	case age == 18:
		fmt.Println("Congratulations! You've just become major!")
		break
	default:
		fmt.Println("You are major!")
		break
	}
}
