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

	var m sync.Mutex

	for i := 0; i < gr; i++ {
		go func() {
			m.Lock()
			n++
			m.Unlock()
			time.Sleep(time.Second / 10)
			wg.Done()
		}()
		go func() {
			defer m.Unlock()
			time.Sleep(time.Second / 10)
			m.Lock()
			n--
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Final value of n = ", n)
}
