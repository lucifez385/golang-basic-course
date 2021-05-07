package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {
}

func main() {
	// resp, err := http.Get("https://www.google.com")
	// if err != nil {
	// 	fmt.Println("Something went wrong!", err)
	// 	os.Exit(1)
	// }

	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Something went wrong!", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	log := logWriter{}
	io.Copy(log, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
