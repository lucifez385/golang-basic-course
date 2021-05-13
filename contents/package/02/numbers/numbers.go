package numbers

import "fmt"

func init() {
	fmt.Println("number init called")
}

func Even(n int) bool {
	return n%2 == 0
}
