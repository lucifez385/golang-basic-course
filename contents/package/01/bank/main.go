package main

import (
	"fmt"
	"mypackages/numbers"
)

func main() {
	n := 10
	fmt.Printf("%d is Even : %t", n, numbers.Even(n))
}
