package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Exercise #1 - int to float and float to int conversion
	var i int = 3
	var f float64 = 3.2
	fmt.Printf("i is %v and type %T\n", float64(i), float64(i))
	fmt.Printf("f is %v and type %T\n", int(f), int(f))

	// Exercise #2 - Conversion
	var ii = 3
	var ff = 3.2
	var s1, s2 = "3.14", "5"
	// 1.int to string
	intToString := strconv.Itoa(ii)
	_ = intToString

	// 2. string to int
	stringToInt, err := strconv.Atoi(s2)
	_, _ = stringToInt, err

	// 3. float to string
	floatToString := fmt.Sprintf("%f", ff)
	_ = floatToString

	// 4.string to float
	stringtoFloat, err := strconv.ParseFloat(s1, 64)
	_, _ = stringtoFloat, err

	// Exercise #3 - Debug
	x, y := 4, 5.1

	z := float64(x) * y
	fmt.Println(z)

	a := 5
	b := 6.2 * float64(a)
	fmt.Println(b)

	// Exercise #4
	const (
		distance = 146_600_000 * 1000
		speed    = 299_792_458
	)
	const time = distance / speed
	fmt.Println(time)

	// Exercise #5 - Logical Operator

	xx, yy := 0.1, 5
	var zz float64

	// Write the correct logical operator (&&, || , !)
	// inside the expression so that result1 will be false and result2 will be true.

	result1 := xx < zz && int(xx) != int(zz)
	fmt.Println(result1)

	result2 := yy == 1*5 || int(zz) == 0
	fmt.Println(result2)
}
