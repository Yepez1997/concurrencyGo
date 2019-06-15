package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	repeatFn := func(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return
				case valueStream <- fn:
				}
			}
		}()
		return valueStream
	}

	take := func(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}

			}
		}()
		return takeStream
	}

	done := make(chan interface{})
	rand.Seed(time.Now().UnixNano())
	randFn := func() interface{} { return rand.Intn(1000) }
	defer close(done)
	for num := range take(done, repeatFn(done, randFn), 10) {
		fmt.Printf("number: %v\n", num)
	}
}
