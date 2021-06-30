package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

// Exercise #1
func cube(param float64) float64 {
	return math.Pow(param, 3)
}

// Exercise #2
func f1(n uint) (int, int) {
	factorial := 1
	sum := 0
	for i := 0; i <= int(n); i++ {
		factorial *= i
		sum += i
	}

	return factorial, sum
}

// Exercise #3
func myFunc(arg string) int {
	i, err := strconv.Atoi(arg)

	if err != nil {
		log.Printf("Cannot convert to int!")
	}
	ii, _ := strconv.Atoi(arg + arg)
	iii, _ := strconv.Atoi(arg + arg + arg)
	return i + ii + iii
}

// Exercise #4
func sum(n ...int) (sum int) { //Exercise #5 naked return
	sum = 0

	for _, v := range n {
		sum += v
	}

	return
}

// Exercise #6
func searchItem(animals []string, search string) bool {
	res := false

	for _, v := range animals {
		if strings.EqualFold(v, search) { // Exercise #7 case - insensitive search
			res = true
		} else {
			res = false
		}
	}

	return res
}

// Exercise #8
func print(msg string) {
	fmt.Println(msg)
}

// Exercise #9
func printInt(n int) {
	fmt.Println(n)
}

func main() {
	// Exercise #1
	cube := cube(5.)
	fmt.Println("Cube:", cube)

	// Exercise #2
	f, s := f1(10)
	fmt.Println("Factorial:", f, "Sum: ", s)

	// Exercise #3
	num := myFunc("5")
	fmt.Printf("Sum: %d\n", num)

	// Exercise #4

	num = sum(5, 10, 15, 20, 30, 40)
	fmt.Printf("Sum: %d\n", num)

	// Exercise #5
	// Naked return
	num = sum(2, 4, 6, 8)
	fmt.Printf("Sum: %d\n", num)

	// Exercise #6
	animals := []string{"lion", "tiger", "bear"}
	result := searchItem(animals, "bear")
	fmt.Println(result) // => true

	// Exercise #7 case-insensitive search
	result = searchItem(animals, "LiOn")
	fmt.Println(result) // => false

	//Exercise #8
	defer print("The Go gopher is the iconic mascot of the Go project.")
	fmt.Println("Hello, Go playground!")

	// Exercise #9
	x := printInt
	fmt.Printf("%T\n", x)
	x(5)
}
