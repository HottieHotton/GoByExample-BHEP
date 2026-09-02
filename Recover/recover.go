package main

import "fmt"

func mayPanic() {
	panic("A problem")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()
	
	//Does not run because mayPanic..well..panics
	fmt.Println("After mayPanic()")
}
