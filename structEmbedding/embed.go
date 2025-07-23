package main

import (
	"fmt"
	"time"
)

// Kind of Inheritance.
type Customer struct {
	name  string
	phone string
}
type Order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	Customer
}

func main() {
	//First Syntax

	/*
		newCustomer := Customer{
			name:  "Ansh",
			phone: "434525",
		}

		Order1 := Order{
			id:       "1",
			amount:   34.22,
			status:   "OK",
			Customer: newCustomer,
		}
		fmt.Println(Order1.Customer)
	*/

	//This is inline syntax.
	Order1 := Order{
		id:     "1",
		amount: 34.22,
		status: "OK",
		Customer: Customer{
			name:  "ansh",
			phone: "4325",
		},
	}
	fmt.Println(Order1.Customer)
	fmt.Println(Order1)
	Order1.Customer.name = "robinHood"
	fmt.Println(Order1)
}
