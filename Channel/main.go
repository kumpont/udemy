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

	ch := make(chan string)
	ch2 := make(chan string)
	ch2 <- "ddd"

	for _, v := range links {
		go checkLink(v, ch)
	}

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-ch)
	// }

	for l := range ch {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, ch)
		}(l)
	}
}

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " is not working!!")
		ch <- link
		return
	}

	fmt.Println(link + " is OK")
	ch <- link
}
