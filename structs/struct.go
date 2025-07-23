package main

import (
	"fmt"
	"time"
)

//order struct

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time //nanoseccond precision(accurate)

}

func (o *order) changeStatus(status string) {
	o.status = status
}
func (o *order) getamount() float32 {
	return o.amount
}
func main() {

	Myorder := order{
		id:     "1",
		amount: 50.00,
		status: "recieved",
	}

	Myorder.changeStatus("confirmed")
	fmt.Println(Myorder)
	fmt.Println(Myorder.getamount())
	//adding fields in structs

	/*
		Myorder.createdAt = time.Now()
		fmt.Println(Myorder)
		fmt.Println(Myorder.id)
	*/

	/*
		Myorder2 := order{
			id:        "2",
			amount:    54.00,
			status:    "Delivered",
			createdAt: time.Now(),
		}
		fmt.Println(Myorder2)
	*/
}
