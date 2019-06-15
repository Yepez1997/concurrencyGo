package main

import (
	"fmt"
	"math/rand"
	"time"
)

// note can often reuse properties and stages of a pipeline multiple times
func main() {
	rand := func() interface{} { return rand.Intn(50000) }

	done := make(chan interface{})

	defer close(done)

	start := time.Now()

	randIntStream := toInt(done, repeatFn(done, rand))
	fmt.Println("Primes: ")
	for prime := range take(done, primeFinder(done, randIntStream), 10) {
		fmt.Printf("\t%d\n", prime)
	}
	fmt.Printf("Search took: %v", time.Since(start))
}
