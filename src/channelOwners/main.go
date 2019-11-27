package main

import "fmt"

func main() {

	chanOwner := func() <-chan int {
		resultStream := make(chan int, 5)
		go func() {
			defer close(resultStream)
			for i := 0; i <= 5; i++ {
				// send the item to the channel
				resultStream <- i
			}
		}()
		return resultStream
	}
	// result stream is a stream of ints
	resultStream := chanOwner()
	// for loop reads from the channel
	for result := range resultStream {
		fmt.Printf("Received %d\n", result)
	}
	fmt.Println("Done receiving")

}
