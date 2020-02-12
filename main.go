package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	// WaitGroup is none of count's concern.
	// Use anonymous function to control group
	go func() {
		count("sheep")
		wg.Done()
	}()

	// wait for it...
	wg.Wait()
}
func count(thing string) {
	for i := 0; i < 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
