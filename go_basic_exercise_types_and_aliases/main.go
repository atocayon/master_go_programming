package main

import "fmt"

func main() {
	// Exercise #1
	type duration = int

	var hour duration
	// a. print out the value of the variable hour
	fmt.Printf("hour is %v\n", hour)
	// b. print out the type of the variable hour
	fmt.Printf("hour is %T\n", hour)
	// c. assign 3600 to the variable hour using the  = operator
	hour = 3600
	// d. print out the value of hour
	fmt.Printf("hour is %v\n", hour)

	// Exercise #2

	var hour2 duration = 3600
	minute := 60
	fmt.Println(hour2 != minute)

	// Exercise #3
	type mile = float64
	type kilomenter = float64

	const m2km = 1.609

	// a.declare a variable called mileBerlinToParis of type mile with value 655.3
	var mileBerlinToParis mile = 655.3
	//b. declare a variable called kmBerlinToParis of type kilometer
	var kmBerlinToParis kilomenter
	//c.calculate the distance between Berlin and Paris in km by multiplying mileBerlinToParis and m2km.
	// Assign the result to kmBerlinToParis
	kmBerlinToParis = mileBerlinToParis * m2km

	fmt.Printf("The distance in km between Berlin to Paris is %v km\n", kmBerlinToParis)
}
