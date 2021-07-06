package main

import "fmt"

func main() {
	// 1. Initialize a float64 channel
	var c1 chan float64

	// 2. Using make()
	// Receive Only
	c2 := make(<-chan rune)
	// Send Only
	c3 := make(chan<- rune)

	// 3.bidirectional buffered channel
	c4 := make(chan int, 10)

	fmt.Printf("%T\n%T\n%T\n%T\n", c1, c2, c3, c4)

}
