package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port int

func SetupServer() *http.Server {
	flag.IntVar(&port, "port", 5000, "port to listen on")
	flag.Parse()

	router := http.NewServeMux()
	router.HandleFunc("/", index)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}
	return &server
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func main() {
	server := SetupServer()
	fmt.Print(server.Addr, "\n")
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("%s", err)
	}
}
