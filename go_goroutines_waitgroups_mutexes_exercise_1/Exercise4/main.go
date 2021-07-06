package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(50)
	// Exercise #4
	for i := 100; i < 150; i++ {
		num := float64(i)
		go func(n float64, wg *sync.WaitGroup) {
			x := math.Sqrt(n)
			fmt.Printf("The square root of %.2f is %.6f\n", n, x)
			wg.Done()
		}(num, &wg)
	}

	wg.Wait()
}
