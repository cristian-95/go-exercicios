package main

import (
	"fmt"
)

/*
1. Modify the program you created in exercise Q2 to use channels, in other words,
the function called in the body should now be a goroutine and communication
should happen via channels. You should not worry yourself on how the goroutine
terminates.
	Q2:
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}*/

func main() {
	c := make(chan int)

	go print(c)
	for i := 0; i < 100; i++ {
		c <- i // sends the value of i into the channel c
	}
}

func print(c chan int) {
	for {
		receptor := <-c // receives the value from channel c and stores into receptor
		fmt.Println(receptor)
	}
}
