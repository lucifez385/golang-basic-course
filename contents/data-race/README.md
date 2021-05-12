# Data Race

## Problem of concurrency

```go
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
```

### Output

```
Final value of n =  -3
```

```
Final value of n =  2
```

```
Final value of n =  0
```


## Race condition

![](01.01.png)

## Race condition detector in Go

```
go run -race main.go
```

## Mutex

```go
var m sync.Mutex
m.Lock()
m.Unlock()
```

