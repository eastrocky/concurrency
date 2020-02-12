package main

import "fmt"

func main() {
	// You can fill up a buffered channel without a receiver and it
	// wont block until the channel is full
	c := make(chan string, 2)
	c <- "hello"
	c <- "world"

	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)
}
