package main

import (
	"fmt"
	"mypackages/numbers"
)

func init() {
	fmt.Println("init called")
}

func init() {
	fmt.Println("init called #2")
}

func main() {
	fmt.Println("main called")
	n := 10
	fmt.Printf("%d is Even : %t", n, numbers.Even(n))
	// init()
}
