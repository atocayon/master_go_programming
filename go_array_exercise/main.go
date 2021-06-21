package main

import "fmt"

func main() {
	// Exercise #1
	var cities [2]string
	grades := [3]float64{1.3, 1.5}
	taskDone := [...]bool{true, false, true, false}
	_ = grades
	_ = taskDone
	fmt.Printf("%#v\n", cities)
	for i := 0; i < len(cities); i++ {
		fmt.Printf("%#v\n", cities[i])
	}

	for _, v := range grades {
		fmt.Printf("%#v\n", v)
	}

	// Exercise #2
	var count int
	nums := [...]int{30, -1, -6, 90, -6}
	for _, v := range nums {
		if v%2 == 0 && v > 0 {
			count++
		}
	}
	fmt.Println("No. of positive intergers in nums:", count)

	// Exercise #3
	myArray := [3]float64{1.2, 5.6}

	myArray[2] = 6

	a := 10
	myArray[0] = float64(a)

	myArray[2] = 10.10

	fmt.Println(myArray)

	// Exercise #4

}
