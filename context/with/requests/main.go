package requests

import (
	"fmt"
	"net/http"
)

func Get(ch chan string, url string) {
	fmt.Println("GET", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Status code:", resp.StatusCode)
	buf := make([]byte, 1024)
	_, err = resp.Body.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	ch <- string(buf[:10]) + "..."
}
