package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1st go routine sleeping")
		time.Sleep(1 * time.Second)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd go routine sleeping")
		time.Sleep(2 * time.Second)
	}()

	wg.Wait()
	fmt.Println("All go routines complete")
}
