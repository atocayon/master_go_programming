package main

import "fmt"

func power(n int, c chan int) {
	c <- n * n
}

func main() {
	c := make(chan int)
	for i := 1; i <= 50; i++ {
		go power(i, c)
	}

	for i := 1; i <= 50; i++ {
		fmt.Println(<-c)
	}

}
