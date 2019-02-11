package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Status code:", resp.StatusCode)

	// // Create byte slice and run it through Read interface to fill it with data
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs) // - Here we feed our byte slice and fill it with data from 'resp.Body.Read'
	// print(string(bs))

	// same thing as above with Read interface
	// io.Copy(os.Stdout, resp.Body) // - Copy 'resp.Body' to 'os.Stdout'

	// same thing as above with Read interface but with 'logWriter' struct and
	// implementing 'Write' method in 'Writer' interface
	lw := logWriter{}
	io.Copy(lw, resp.Body) // - Copy 'resp.Body' to 'lw', logWriter implements 'Writer' interface with 'Write' method

}

func (logWriter) Write(bs []byte) (int, error) {

	fmt.Println(string(bs))
	fmt.Println("\n Bytes cound: ", len(bs))
	return len(bs), nil
}
