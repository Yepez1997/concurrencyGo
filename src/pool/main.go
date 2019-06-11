package main

import (
	"fmt"
	"sync"
)

func main() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating a new instance")
			return struct{}{}
		},
	}

	// creates a new instance at first
	myPool.Get()
	instance := myPool.Get()
	myPool.Put(instance)
	myPool.Get()

}
