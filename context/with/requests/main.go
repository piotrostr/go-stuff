package requests

import (
	"net/http"
	"strings"
	"time"
)

func Get(ch chan string, url string) {
	println("GET", strings.Split(url, "//")[1])
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		println(err)
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 1024)
	_, err = resp.Body.Read(buf)
	if err != nil {
		println(err)
		return
	}
	elapsed := time.Since(start)
	ch <- string(url + ": " + elapsed.String())
}
