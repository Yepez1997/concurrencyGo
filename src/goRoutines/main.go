package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("hello")
}

func main() {

	go sayHello()

	sayHello2 := func() {
		fmt.Println("hello from the other side")
	}

	go sayHello2()

	go func() {
		fmt.Println("hello from anonymous")
	}()

	time.Sleep(3 * time.Second)

}
