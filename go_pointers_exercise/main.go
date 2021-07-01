package main

import "fmt"

// Exercise #3
func swap(x *float64, y *float64) {
	*x, *y = *y, *x
}

func main() {
	// Exercise #1
	x := 10.10

	//Print out the address of x
	fmt.Printf("The address of x is %p\n", &x)

	// Declare a pointer and stores the address
	ptr := &x

	//Print out the value and type of ptr
	fmt.Printf("Type of ptr: %T Value of ptr: %v\n", ptr, ptr)

	//Print out the address of pointer and the value of x through the pointer
	fmt.Printf("The address of the pointer is %p, and the value of x  through the pointeris %v\n", &ptr, *ptr)

	// Exercise #2
	x2, y2 := 10, 2
	ptrx, ptry := &x2, &y2
	z := *ptrx / *ptry
	fmt.Printf("z is %d\n", z)

	// Exercise #3
	x3, y3 := 5.5, 8.8
	fmt.Printf("The value of x is %v, the value of y is %v before swapping\n", x3, y3)
	swap(&x3, &y3)
	fmt.Printf("The value of x is %v, the value of y is %v after swapping\n", x3, y3)

}
