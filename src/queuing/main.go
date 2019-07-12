package main

import "time"

// queing is used to reduce the time in the blocking state
// decouple stages such that the runtime of one does not affect the runtime of another
func main() {
	done := make(chan interface{})
	defer close(done)

	zeros := take(done, 3, repeat(done, 0))
	short := sleep(done, 1*time.Second, zeroes)
	// buffer := buffer(done, 2, short)
	// if added a buffer shorts stage will complete in 3 seconds rather than 9
	// although the total time remains at 13 seconds
	long := sleep(done, 4*time.Second, short)
	pipeline := long
	// building pipelines in go

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

}
