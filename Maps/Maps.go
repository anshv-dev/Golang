package main

import (
	"fmt"
	"maps"
)

// maps->hash,Object
func main() {
	/*
		m := make(map[string]string) //map syntax

		//setting an element
		m["name"] = "Golang"
		m["area"] = "backend"
	*/

	//Imp:If key value does not exist then it will return zero value.
	// fmt.Println(m["pj"])

	/*
		a := make(map[string]int)
		a["age"] = 90
		fmt.Println(len(m))
		fmt.Println(len(a))
		delete(a, "age") //delete syntax
		a["age"] = 45
		a["price"] = 3000
		a["NumberOfLetters"] = 34
		fmt.Println(a)

		//if there is need to wipe out the elements
		clear(a)
		fmt.Println(a)

	*/

	//Another way to create map
	/*
		m := map[string]int{
			"price":  34,
			"age":    34,
			"salary": 67455,
		}
		fmt.Println(m)
	*/

	/*
		m := map[string]int{"name": 23, "age": 34}

		//b will get the value.
		//Ok will get true or false according to the situation.
		b, Ok := m["name"]
		// fmt.Println(a)
		fmt.Println(b)  //Output:23
		fmt.Println(Ok) //Output:true
	*/

	m1 := map[string]int{"name": 23, "age": 34}
	m2 := map[string]int{"name": 23, "age": 334}
	fmt.Println(maps.Equal(m1, m2))
}
