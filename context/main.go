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
	requests.Get(GOOGLE)
	requests.Get(YAHOO)
	requests.Get(BING)
}
