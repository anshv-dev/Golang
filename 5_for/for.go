package main

import "fmt"

func main() {
	//while-loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	//infinite-loop
	for {
		println("1")
	}

	//classic-for-loop
	for i := 0; i <= 3; i++ {
		fmt.Println(i)
	}

	//range in loop
	for i := range 7 {
		fmt.Println(i)
	}
}
