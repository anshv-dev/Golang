package main

import "fmt"

func main() {
	//if array type is int then default value is 0,if string then "",if bool then false.

	// var nums [4]int
	// nums[0] = 1
	// nums[1] = 2
	// fmt.Println(nums[0])
	// fmt.Println(len(nums))
	// fmt.Println(nums)

	// var vals [4]bool
	// fmt.Println(vals)

	// var val [4]string
	// fmt.Println(val)

	//to declare it in single line

	// nums := [3]int{1, 2, 3}
	// fmt.Println(nums)

	//2-D marrays

	numms := [2][2]int{{1, 2}, {23, 3}}

	fmt.Println(numms)

}
