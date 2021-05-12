package main

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()
		getAndPrintData("https://qvault.io")
	}()
	go func() {
		defer wg.Done()
		getAndPrintData("https://github.com")
	}()
	go func() {
		defer wg.Done()
		getAndPrintData("https://gitlab.io")
	}()
	wg.Wait()

	fmt.Println("nextProcess >>")
}

func getAndPrintData(url string) {
	resp, _ := http.Get(url)
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	fmt.Println(buf.String())
}
