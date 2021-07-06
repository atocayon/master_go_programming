package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	// Exercise #3
	go func(n float64, wg *sync.WaitGroup) {
		x := math.Sqrt(n)
		fmt.Printf("The square root of %.2f is %.6f\n", n, x)
		wg.Done()
	}(5., &wg)

	wg.Wait()
}
