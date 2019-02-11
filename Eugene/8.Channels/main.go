package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		// fmt.Println(link)
		go checkLink(link, c)
	}

	// Finite
	// for i := 0; i < len(links); i++ {
	// 	// fmt.Printf(<-c) // - receive data from channel, and print it
	// 	go checkLink(<-c, c)
	// }

	// // Infinite
	// for {
	// 	// fmt.Printf(<-c) // - receive data from channel, and print it
	// 	go checkLink(<-c, c) // - receive data from channel, and execute it again with returned data in '<-c'
	// }

	// Infinite, same syntax as Infinite above
	// for l := range c {
	// 	go checkLink(l, c) // - receive data from channel, and execute it again with returned data in '<-c'
	// }

	//// Infinite With sleep
	for l := range c {
		// we put sleep code inside go routine, we are not blocking 'Main' routine
		// FUNCTION Literal in Golang (lambda in python, anonymous function in javascript)
		go func(link string) { // link passed from Literal (l)
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("might be down: ", link)
		c <- link // - send data to channel
		return
	}
	fmt.Println("is up: ", link)
	c <- link // - send data to channel
}
