package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	c := make(chan int)

	go func() {
		time.Sleep(5 * time.Second)
		close(c)
	}()

	fmt.Println("Blocking on read...")
	select {
	case <-c:
		fmt.Printf("Unblocking %v later .\n", time.Since(start))
	}
}
