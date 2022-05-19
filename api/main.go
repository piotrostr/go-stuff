package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func runServer() {
	port := 8000
	host := "0.0.0.0:"
	http.HandleFunc("/", homePage)
	fmt.Printf("listening on localhost:%d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s%d", host, port), nil))
}

func main() {
	runServer()
}
