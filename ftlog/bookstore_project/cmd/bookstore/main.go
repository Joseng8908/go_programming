package main

import (
	"fmt"
)

func main() {
	var title string
	var name string
	var copies int
	var edition string
	var special_offer bool
	var discount_percenage float64
	title = "For the Love of Go"
	name = "Shakespeare"
	copies = 99
	edition = "First"
	special_offer = true
	discount_percenage = 0.1
	fmt.Println(title)
	fmt.Println(copies)
	fmt.Println(name)
	fmt.Println(edition)
	fmt.Println(special_offer)
	fmt.Println(discount_percenage)
}
