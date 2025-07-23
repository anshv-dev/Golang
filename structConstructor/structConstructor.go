package main

import (
	"fmt"
	"time"
)

type Order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
}

// Creating something like constructor.
func newOrder(id string, amount float32, status string) *Order {
	myOrder := Order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &myOrder
}
func main() {
	myOrder := newOrder("1", 30, "recieved")
	fmt.Println(myOrder)
	fmt.Println(myOrder.amount)

	//Example of Inline struct.
	language := struct {
		name   string
		isGood bool
	}{"golang", true}
	fmt.Println(language)
}
