package main

import "fmt"

func main() {

	// Exercise #4
	type Grades struct {
		grade  float64
		course string
	}

	// Exercise #1
	type Person struct {
		name      string
		age       int
		fav_color []string
		grades    Grades
	}

	me := Person{name: "aljon", age: 24, fav_color: []string{"black", "gray"}, grades: Grades{grade: 80, course: "Programming"}}   // grades fields for exercise #4
	you := Person{name: "kira", age: 25, fav_color: []string{"red", "green", "blue"}, grades: Grades{grade: 83, course: "Golang"}} //grades fields for exercise #4

	fmt.Printf("%#v\n", me)
	fmt.Printf("%#v\n", you)

	// Exercise #2
	me.name = "Andrei"

	var colors []string = you.fav_color
	fmt.Printf("Color type: %T Color value: %#v", colors, colors)

	colors = append(colors, "violet")
	you.fav_color = colors

	fmt.Printf("%#v\n", me)
	fmt.Printf("%#v\n", you)

	// Exercise #3

	for _, v := range me.fav_color {
		fmt.Printf("My favorite color is: %#v", v)
	}

	// Exercise #4
	me.grades.grade = 98
	me.grades.course = "Golang"

	fmt.Printf("%#v\n", me)
	fmt.Printf("%#v\n", you)
}
