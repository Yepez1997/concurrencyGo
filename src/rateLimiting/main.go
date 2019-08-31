package main

import (
	"context"
	"log"
	"os"
	"sync"
)

func Open() *APIConnection {
	return &APIConnection{}
}

type APIConnection struct{}

func (a *APIConnection) ReadFile(ctx context.Context) error {
	return nil
}

func (a *APIConnection) ResolveAdress(ctx context.Context) error {
	return nil
}

// main function calls
func main() {
	defer log.Printf("Done.")
	// set up the output
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)
	apiConnection := Open()

	// 20 requests
	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := apiConnection.ReadFile(context.Background())
			if err != nil {
				log.Printf("cannot ReadFile: %v", err)
			}
			log.Printf("Readfile")
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			err := apiConnection.ResolveAdress(context.Background())
			if err != nil {
				log.Printf("cannot Resolve Address: %v", err)
			}
			log.Printf("ResolveAddress")
		}()
	}
	// wait til all the go routines ran for the wait group
	wg.Wait()

}
