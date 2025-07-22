package main

import "fmt"

/*
func add(a, b int) int {
	return a + b
}*/

func add(a int, b int) int { //single return function
	return a + b
}

func retuString() (string, string, string) {
	return "golang", "c++", "javascript"
}
func process() func(a int) int {
	return func(a int) int {
		return 4
	}
}
func main() {
	//add function
	result := add(3, 5)
	fmt.Println(result)

	//string multiple return function
	fmt.Println(retuString())
	//distribute the values
	lang1, lang2, lang3 := retuString()
	fmt.Println(lang1)
	fmt.Println(lang2)
	fmt.Println(lang3)

	fn := process()
	fmt.Println(fn(6))

}
