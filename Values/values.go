package main

import "fmt"

func main() {
	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)


	// Personal TESTING Println behavior

	fmt.Println("1 {}", 1, 2, 3)
	//Result = 1 {} 1 2 3

	fmt.Println(4)
	//Result = 4

	fmt.Println("5\n5\n5")
	//Result = 5 5 5 on 3 lines. Behaves the same on Python

	fmt.Println("{\"name\": \"John\", \"age\": 30, \"city\": \"New York\"}")
	//Result = {"name": "John", "age": 30, "city": "New York"}
}
