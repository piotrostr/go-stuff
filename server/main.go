package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var port int

func SetupServer() *http.Server {
	flag.IntVar(&port, "port", 5000, "port to listen on")
	flag.Parse()

	router := http.NewServeMux()
	router.HandleFunc("/", index)

	logger := log.New(os.Stdout, "", log.LstdFlags)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  5 * time.Second,
	}
	return server
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

func main() {
	server := SetupServer()

	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		server.ErrorLog.Println("shutting down...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			server.ErrorLog.Fatalf("could not gracefully shutdown %s", err)
		}
		close(done)
	}()

	server.ErrorLog.Printf("starting on %s\n", server.Addr)

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("%s", err)
	}

	<-done
	server.ErrorLog.Println("server stopped")
}
