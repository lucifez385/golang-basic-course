package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const gr = 100

	var wg sync.WaitGroup
	wg.Add(gr * 2)

	var n int = 0

	for i := 0; i < gr; i++ {
		go func() {
			defer wg.Done()
			n++
			time.Sleep(time.Second / 10)
		}()
		go func() {
			defer wg.Done()
			n--
			time.Sleep(time.Second / 10)
		}()
	}

	wg.Wait()

	fmt.Println("Final value of n = ", n)
}
