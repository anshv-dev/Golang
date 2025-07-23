package main

import "fmt"

//by value will not change source value.
/*
func change( num int ) {
	num = 5
	fmt.Println(num)
}
*/

//Pointer Concept
func changepointer(num *int) {
	*num = 5
	fmt.Println(*num)
}
func main() {
	num := 1
	// change(num)
	changepointer(&num)
	fmt.Println(num)
}
