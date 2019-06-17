package main

import "time"

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
}
