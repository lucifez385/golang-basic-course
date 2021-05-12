package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)

	fmt.Println("Channel's capacity:", cap(c1))

	go func(c chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Printf("func goroutine #%d starts sending data into the channel\n", i)
			c <- i
			fmt.Printf("func goroutine #%d after sending data into the channel\n", i)
		}
		close(c)

	}(c1)

	fmt.Println("main goroutine sleeps 2 seconds")
	time.Sleep(time.Second * 2)

	for v := range c1 {
		fmt.Println("main goroutine received value from channel:", v)
	}

	fmt.Println(<-c1) // => 0

	// c1 <- 10 // => panic: send on closed channel
}
