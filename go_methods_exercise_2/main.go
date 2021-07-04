package main

import "fmt"

// Exercise #1
type vehicle interface {
	License() string
	Name() string
}

// Exercise #3
type cube struct {
	edge float64
}

type car struct {
	licenceNo string
	brand     string
}

func (c car) License() string {
	return c.licenceNo
}

func (c car) Name() string {
	return c.brand
}

// Exercise #3
func volume(c cube) float64 {
	return c.edge * c.edge * c.edge
}

func main() {
	// Exercise #1
	var v vehicle = car{licenceNo: "1234", brand: "Wigo gen 2"}

	fmt.Println(v.License())
	fmt.Println(v.Name())

	// Exercise #2
	var empty interface{}
	fmt.Printf("%T\n", empty) // <nil>
	empty = 1
	fmt.Printf("%T\n", empty) // int
	empty = 1.1
	fmt.Printf("%T\n", empty) // float64
	empty = []int{}
	fmt.Printf("%T\n", empty) // []int
	// 5. Type Assertion
	empty = append(empty.([]int), 10)

	// 6.
	fmt.Printf("%v\n", empty)

	// Exercise #3
	var x interface{}
	x = cube{edge: 5.}
	// Type Assertion
	c := volume(x.(cube))
	fmt.Printf("Cube Volume: %v\n", c)
}
