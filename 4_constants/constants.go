package main

import "fmt"

const age = 90

func main() {
	const name = "ANSH VERMA"

	fmt.Println(age)
	fmt.Println(name)

	const (
		port = 3000
		host = "localhost"
		tech = "tech giant"
	)

	fmt.Println(port, host, tech)
}
