package main

import "fmt"

func main() {
	c := make(chan int)
	for i := 1; i <= 50; i++ {
		go func(n int) {
			c <- n * n
		}(i)
	}

	for i := 1; i <= 50; i++ {
		fmt.Println(<-c)
	}

}
