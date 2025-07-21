package main

import "fmt"

//Iterating over data structure
func main() {
	/*
		nums := []int{6, 7, 8}

			for i := 0; i < len(nums); i++ {
				fmt.Println(nums[i])
			}
	*/

	//Another syntax
	/*
		sum := 0
		for _, num := range nums {
			sum += num
		}
		fmt.Println(sum) //Output:=21
	*/

	/*
		for i, num := range nums {
			fmt.Println(i, num)
		}
	*/

	/*
		m := map[string]string{"fname": "john", "lname": "doe"}
		for k, v := range m {
			fmt.Println(k, v)
		}
	*/

	//this will print unicode of each character
	for i, c := range "gol/" {
		fmt.Println(i, string(c))
	}
}
