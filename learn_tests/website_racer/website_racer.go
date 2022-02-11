package website_racer

import (
	"errors"
	"net/http"
	"time"
)

var TimeoutError = errors.New("timeout")

func Racer(url1 string, url2 string) (winner string, err error) {
	select {

	case <-ping(url1):
		return url1, nil

	case <-ping(url2):
		return url2, nil

	case <-time.After(10 * time.Second):
		return "", TimeoutError
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
