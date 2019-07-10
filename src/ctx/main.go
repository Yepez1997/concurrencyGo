package main

import (
	"fmt"
	"sync"
	"time"
)

// with done channel only 

func main() {
	var wg sync.WaitGroup
	done := make(chan interface{})
	// cloe the chanel when done
	defer close(done)

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printGreeting(done); err != nil {
			// print the error
			fmt.Printf("%v", err)
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := printFarewell(done); err != nil {
			fmt.Printf("%v", err)
			return
		}
	}()

	wg.Wait()

}

func printGreeting(done <-chan interface{}) error {
	greeting, err := genGreeting(done)
	if err != nil {
		return nil
	}
	fmt.Printf("%s world", greeting)
	return nil
}

func printFarewell(done <-chan interface{}) error {
	farewell, err := genFarewell(done)
	if err != nil {
		return err
	}
	fmt.Printf("%s world!\n", farewell)
	return nil
}

func genGreeting(done <-chan interface{}) (string, error) {
	switch locale, err := local(done); {
	case err != nil:
		return "", err

	case locale == "EN/US":
		return "hello", nil
	}
	return "", fmt.Errorf("unsupported locale")
}

func genFarewell(done <-chan interface{}) (string, error) {
	switch locale, err := locale(done); {
		case err != nil:
			return  "", err
		case locale == "EN/US"
			return "hello", nil 
	}
	return "", fmt.Errorf("unsupported locale")
}

func local(done <-chan interface{}) (string, error) {
	select {
	case <-done:
		return "", fmt.Errorf("cancelled")

	case <-time.After(1 * time.Minute):
	}
	return "EN/US", nil
}

// should print good bye world 
// then followed by hello world 
// two branchges of the program are set to run concurrently 
