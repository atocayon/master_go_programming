package main

import "fmt"

func main() {
	// Exercise #1
	// const daysWeek int32 = 7
	// const lightSpeed int64 = 299792458
	// const pi float64 = 3.14159

	// Exercise #2
	const (
		daysWeek   = 7
		lightSpeed = 299792458
		pi         = 3.14159
	)

	// Exercise #3
	const secPerDay int64 = 86400
	const daysYear int64 = 365
	fmt.Printf("The total number of seconds in a year is %v\n", secPerDay*daysYear)

	// Exercise #4
	const x int = 10

	// declaring a constant of type slice int ([]int)
	// const m = []int{1: 3, 4: 5, 6: 8} -> You cannot declare a slice constant.

	// Exercise #5
	const a int = 7
	const b float64 = 5.6
	const c = float64(a) * b

	xx := 8
	_ = xx
	// const xc int = xx // variables belong to runtime

	// const noIPv6 = math.Pow(2, 128) // functions calls belong to runtime

	// Exercise #6

	const (
		Jun = iota + 6
		Jul
		Aug
	)

	fmt.Println(Jun, Jul, Aug)

}
