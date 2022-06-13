package main

import (
	"ctx/with/requests"
	_ "ctx/with/server"
)

const (
	GOOGLE = "https://www.google.com"
	YAHOO  = "https://www.yahoo.com"
	BING   = "https://www.bing.com"
)

func main() {
	google := make(chan string)
	yahoo := make(chan string)
	bing := make(chan string)
	go requests.Get(google, GOOGLE)
	go requests.Get(yahoo, YAHOO)
	go requests.Get(bing, BING)
	for i := 0; i < 3; i++ {
		select {
		case s := <-google:
			println(s)
		case s := <-yahoo:
			println(s)
		case s := <-bing:
			println(s)
		}
	}
}
