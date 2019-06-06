package main

import (
	"fmt"
	"sync"
)

func main() {
	type Button struct {
		Clicked *sync.Cond
	}
	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	subscribe := func(c *sync.Cond, fn func()) {
		var goRoutineRunning sync.WaitGroup
		goRoutineRunning.Add(1)
		go func() {
			goRoutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		goRoutineRunning.Wait()
	}

	var ClickedRegistered sync.WaitGroup
	ClickedRegistered.Add(3)
	subscribe(button.Clicked, func() {
		fmt.Println("Maximizing Window")
		ClickedRegistered.Done()
	})

	subscribe(button.Clicked, func() {
		fmt.Println("Displaying annoying anaolog box")
		ClickedRegistered.Done()
	})

	subscribe(button.Clicked, func() {
		fmt.Println("Mouse Clicked")
		ClickedRegistered.Done()
	})

	button.Clicked.Broadcast()
	ClickedRegistered.Wait()
}
