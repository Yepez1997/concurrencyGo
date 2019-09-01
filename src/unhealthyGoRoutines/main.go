package main

import "time"

func main() {
	type startGoRoutineFn func(
		done <-chan interface{},
		pulseInterval time.Duration,
	) (hearbeat <-chan interface{})

	newSteward := func(
		timeout time.Duration,
		startGoRoutine starstartGoRoutineFn,
	) starstartGoRoutineFn {
		return func(
			done <-chan interface{}
			pulseInterval time.Duration
		) (<-chan interface{}) {
			hearbeat := make(chan interface{})
			go func() {
				defer close(heartbeat)
				var wardDone chan interface{}
				var wardHeartBeat <-chan interface{}
				startWard := func() {
					wardDone = make(chan interface{})
					wardHeartBeat = startGoRoutine(or(warDone, done), timeout/2)
				}
				
				startWard()
				pulse := time.Tick(pulseInterval)
				monitorLoop:
					for {
						timeoutSignal := time.After(timeout)
						for {
							select {
							case <- pulse:
								select {
								case heatbeat <- struct{}{}:
								default:
								}
							case <- wardHeartBeat:
								continue monitoLoop
							case <- timeoutSignal:
								log.Println("Steward: war unhealthy ... restarting ")
								close(wardDone)
								startWard()
								continue monitorLoop
							case <-done:
								return
							}
						}
					}
			}()
			return hearbeat
		}
	}
}
