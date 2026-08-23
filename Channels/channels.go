package main

import "fmt"

func main() {
	messages := make(chan string)
	testerMath := make(chan int)
	testfunc := make(chan int)

	go func() { messages <- "ping"}()
	go func() { testerMath <- 2+2}()
	go func() { testfunc <- func(a, b int) int { return a / b } (2, 2) }()
	 
	msg := <-messages
	tst := <-testerMath
	tstfunc := <-testfunc
	fmt.Println(msg)
	fmt.Println(tst)
	fmt.Println(tstfunc)
}