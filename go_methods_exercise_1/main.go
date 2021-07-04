package main

import "fmt"

type money float64 //Exercise #1

// Exercise #3
type book struct {
	price float64
	title string
}

// Exercise #1
func (m money) print() {
	fmt.Printf("%.2f\n", m)
}

// Exercise #2
func (m money) printStr() string {
	myStr := fmt.Sprintf("%.2f", m) //Convert float to string
	return myStr
}

// Exercise #3
func (b book) vat() float64 {
	return b.price * 0.09
}

func (b *book) discount() {
	b.price = b.price * 0.09
}

// Exercise #4
// method for book type
func (b *book) changePrice(p float64) {
	b.price = p
}

func main() {
	// Exercise #1
	var salary money = 1.5 * 5.503
	fmt.Println(salary) // 8.2545

	salary.print() //8.25

	// Exercise #2
	myStr := salary.printStr()
	fmt.Printf("Type: %T Value: %s\n", myStr, myStr)

	//Exercise #3
	bestBook := book{title: "The Trial", price: 9.9}

	// calling the methods
	vat := bestBook.vat()
	fmt.Printf("Vat:%v\n", vat)

	bestBook.discount()
	fmt.Printf("%#v\n", bestBook)

	// Exercise #4
	// changing the price
	bestBook.changePrice(11.99)

	fmt.Printf("Book's price:%.2f\n", bestBook.price)
}
