package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()
	go func() {
		for {
			c2 <- "every 2 seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		fmt.Println(<-c1)
		// Receiving is a blocking call
		fmt.Println(<-c2)
	}
}
