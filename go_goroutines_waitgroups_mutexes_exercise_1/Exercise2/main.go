package main

import (
	"fmt"
	"sync"
)

// Exercise #2

func sum(n1, n2 float64, wg *sync.WaitGroup) {

	fmt.Printf("Sum: %f\n", n1+n2)

	wg.Done()

}

func main() {

	var wg sync.WaitGroup

	wg.Add(3)
	// Exercise #2
	go sum(2.5, 3.5, &wg)
	go sum(1.2, 5., &wg)
	go sum(5., 4.3, &wg)

	wg.Wait()

}
