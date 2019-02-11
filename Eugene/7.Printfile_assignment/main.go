package main

import (
	"fmt"
	"io"
	"os"
)

// go run main.go myfile.txt

func main() {
	filename, error := os.Open(os.Args[1])
	if error != nil {
		fmt.Println("Error:", error)
		os.Exit(1)
	}
	io.Copy(os.Stdout, filename)
}
