package main

import (
	"fmt"
	"reflect"
)

func counter() func() int {
	var count int = 0
	return func() int {
		count += 1
		return count
	}
}
func main() {
	fmt.Println(reflect.TypeOf(counter()))
	increment := counter()
	fmt.Println(reflect.TypeOf(increment()))
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}
