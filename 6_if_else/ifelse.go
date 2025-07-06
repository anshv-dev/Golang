package main

import "fmt"

func main() {
	age := 14
	if age >= 18 {
		fmt.Println("MAha Adult")
	} else if age < 18 {
		fmt.Println("teenager")
	} else {
		fmt.Println("balak")
	}

	//we can declare a variable inside if construct

	if age := 15; age >= 18 {
		fmt.Println("person is an adult", age)
	} else if age >= 12 {
		fmt.Println("person is teenager", age)
	}

	//go does not have ternary operator, we have to use normal if else

}
