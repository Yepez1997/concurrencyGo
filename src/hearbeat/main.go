package main

import (
	"fmt"
	"time"
)

// function to send heartbeats
func main() {
	eventEmit()
}

func eventEmit() {
	doWork := func(done <-chan interface{}, pulseInterval time.Duration) (<-chan interface{}, <-chan time.Time) {
		heartbeat := make(chan interface{})
		results := make(chan time.Time)
		go func() {
			defer close(heartbeat)
			defer close(results)
			pulse := time.Tick(pulseInterval)
			workGen := time.Tick(2 * pulseInterval)

			// actual heatbeat implementation
			sendPulse := func() {
				select {
				// send a heartbeat to the channel
				case heartbeat <- struct{}{}:
				default:
				}
			}

			sendResult := func(r time.Time) {
				for {
					select {
					case <-done:
						return
					case <-pulse:
						sendPulse()
						// once you send the result terminate the loop
					case results <- r:
						return
					}
				}
			}
			for {
				// keeps looping forever until it is done
				select {
				case <-done:
					return
					// waiting to receive a pulse
				case <-pulse:
					// once the pulse is received send a pulse
					sendPulse()
				case r := <-workGen:
					// send the result of the worker
					sendResult(r)
				}
			}
		}()
		return heartbeat, results
	}

	done := make(chan interface{})
	time.AfterFunc(10*time.Second, func() {
		close(done)
	})
	const timeout = 2 * time.Second
	heartbeat, results := doWork(done, timeout/2)
	for {
		select {
		case _, ok := <-heartbeat:
			if ok == false {
				return
			}
			fmt.Println("Pulse")
		case r, ok := <-results:
			if ok == false {
				return
			}
			fmt.Printf("Results %v\n", r.Second())
		case <-time.After(timeout):
			return
		}
	}
}
