package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	port := 6969
	host := "localhost:"
	http.HandleFunc("/", homePage)
	fmt.Printf("listening on localhost:%d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s%d", host, port), nil))
}

func main() {
	handleRequests()
}
