package main

import "fmt"

func main() {

	// Exercise #1
	x, y, z := 10, 15.5, "Gophers"
	score := []int{10, 20, 30}

	// 1.
	fmt.Printf("x is %d\n", x)          // decimal
	fmt.Printf("y is %f\n", y)          // float
	fmt.Printf("z is %s\n", z)          // string
	fmt.Printf("score is %#v\n", score) // slice

	// 2.
	fmt.Printf("z is %q\n", z) // quoted string

	// 3.
	fmt.Printf("x is %v\n", x)
	fmt.Printf("y is %v\n", y)
	fmt.Printf("z is %v\n", z)
	fmt.Printf("score is %v\n", score)

	// 4.
	fmt.Printf("y type is %T\n", y)
	fmt.Printf("score type is %T\n", score)

	// Exercise #2
	const xx float64 = 1.422349587101

	fmt.Printf("x is %.4f\n", xx) // -> Prints 4 decimal points

	// Exercise #3
	shape := "circle"
	radius := 3.2
	const pi float64 = 3.14159
	circumference := 2 * pi * radius

	fmt.Printf("Shape: %q\n", shape)
	fmt.Printf("Circle's circumference with radius %f is %f\n", radius, circumference)
	_ = shape
}
