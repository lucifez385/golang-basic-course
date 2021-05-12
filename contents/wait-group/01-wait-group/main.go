package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	go getAndPrintData("https://qvault.io")
	go getAndPrintData("https://github.com")
	go getAndPrintData("https://gitlab.io")
	fmt.Println("nextProcess >>")
}

func getAndPrintData(url string) {
	resp, _ := http.Get(url)
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	fmt.Println(buf.String())
}
