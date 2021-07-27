package main

import "fmt"

var (
	name string
	age int
)

func main() {
	fmt.Println("Your name:")
	fmt.Scan(&name)
	fmt.Println("Age:")
	fmt.Scan(&age)
	fmt.Printf("Name: %s, age: %d\n", name, age)
}
