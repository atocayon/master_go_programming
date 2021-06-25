package main

import "fmt"

func main() {
	// Exercise #1
	var name = "Aljon"
	country := "Philippines"
	fmt.Printf("Your name: %s\nCountry: %s\n", name, country)
	fmt.Println(`a) He says "Hello"`)
	fmt.Println(`b) C:\Users\a.txt`)

	// Exercise #2
	r := 'ă'
	fmt.Printf("The type of r: %T\n", r) // rune is alias to int32

	s1, s2 := "ma", "m"

	s := s1 + s2 + string(r) // converting rune to string (expliction conversion is required)
	fmt.Printf("String and rune concatenation: %s", s)

	// Exercise #3
	s1 = "țară means country in Romanian"

	fmt.Println("Byte in string:")
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%v\n", s1[i])
	}

	for i, r := range s1 {
		fmt.Printf("byte index: %d -> rune: %c\n", i, r)
	}

	// Exercise #4

	s1 = "Go is cool!"
	x := s1[0]
	fmt.Println(x)

	// s1[0] = 'x' // error string is immutable

	// printing the number of runes in the string
	fmt.Println(len(s1))

	// Exercise #5

	s = "你好 Go!"
	b := []byte(s)

	fmt.Printf("%#v\n", b)

	for i, r := range b {
		fmt.Printf("%d -> %d\n", i, r)
	}

	// Exercise #6

	s = "你好 Go!"
	rune := []rune(s)

	fmt.Printf("%#v\n", rune)

	for i, r := range rune {
		fmt.Printf("%d -> %d\n", i, r)
	}
}
