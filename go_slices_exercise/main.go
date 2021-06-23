package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	separator := strings.Repeat("#", 20)
	// Exercise #1
	names := []string{"Arnold", "Mark", "Gabriel"}
	for i, v := range names {
		fmt.Printf("index: %d value: %s \n", i, v)
	}

	fmt.Println(separator)
	// Exercise #2
	mySlice := []float64{1.2, 5.6}

	// mySlice[2] = 6 - index out of range

	a := 10
	mySlice[0] = float64(a)

	// mySlice[3] = 10.10 - index out of range

	mySlice = append(mySlice, float64(a))

	fmt.Println(mySlice)

	fmt.Println(separator)
	// Exercise #3

	nums := []float64{1.1, 1.2, 1.3}
	nums = append(nums, 10.1)
	nums = append(nums, 4.1, 5.5, 6.6)

	fmt.Printf("nums slice: %v\n", nums)
	n := []float64{1.4, 1.5}
	nums = append(nums, n...)

	fmt.Printf("nums slice: %v\n", nums)

	fmt.Println(separator)
	// Exercise #4

	if len(os.Args) < 2 { //if not run with arguments
		log.Fatal("Please run the program with arguments (2-10 numbers)!")

	}

	//taking the arguments in a new slice
	args := os.Args[1:]

	// declaring and initializing sum and product of type float64
	sum, product := 0., 1.

	if len(args) < 2 || len(args) > 10 {
		fmt.Println("Please enter between 2 and 10 numbers!")
	} else {

		for _, v := range args {
			num, err := strconv.ParseFloat(v, 64)
			if err != nil {
				continue //if it can't convert string to float64, continue with the next number
			}
			sum += num
			product *= num
		}
	}

	fmt.Printf("Sum: %v, Product: %v\n", sum, product)

	fmt.Println(separator)
	// Exercise #5

	nums5 := []int{5, -1, 9, 10, 1100, 6, -1, 6}
	nums5 = nums5[1 : len(nums5)-2]
	nums5_sum := 0
	for _, v := range nums5 {
		nums5_sum += v
	}

	fmt.Printf("sum: %#v\n", nums5_sum)

	fmt.Println(separator)
	// Exercise #6

	friends := []string{"Marry", "John", "Paul", "Diana"}
	friendsCopy := make([]string, len(friends))
	copy(friendsCopy, friends) //using copy()

	fmt.Printf("Friends: %v\n", friends)                       //Original Friends
	fmt.Printf("Friends Copy using copy(): %v\n", friendsCopy) // Friends Copy

	friendsCopy[0] = "Aljon" //Modifying friends copy

	fmt.Printf("Friends: %v\n", friends)                        //Original Friends
	fmt.Printf("Friends Copy using copy() : %v\n", friendsCopy) // Friends Copy

	fmt.Println(separator)
	// Exercise #7

	friendsCopy2 := []string{}
	friendsCopy2 = append(friends, friendsCopy2...)

	fmt.Printf("Friends: %v\n", friends)                           //Original Friends
	fmt.Printf("Friends Copy using append() : %v\n", friendsCopy2) // Friends Copy

	friendsCopy2[0] = "Aljon"

	fmt.Printf("Friends: %v\n", friends)                           //Original Friends
	fmt.Printf("Friends Copy using append() : %v\n", friendsCopy2) // Friends Copy

	fmt.Println(separator)
	// Exercise #8
	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	newYears := []int{}

	newYears = append(years[:3], years[len(years)-3:]...)
	fmt.Printf("New Years: %#v\n", newYears)

}
