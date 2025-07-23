package main

import "fmt"

func sum(x ...int) int {
	total := 0
	for _, num := range x {
		total += num
	}
	return total
}

func sum1(x ...int) int {
	total := 0
	for _, num := range x {
		total += num
	}
	return total
}
func main() {

	fmt.Println(sum(1, 2, 3, 4))

	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(sum1(nums...))
}
