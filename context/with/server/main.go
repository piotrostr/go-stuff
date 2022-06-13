package server

import (
	"fmt"
	"net/http"
	"time"
)

func Run() {
	http.ListenAndServe(":8080", http.HandlerFunc(handler))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	select {
	case <-time.After(2 * time.Second):
		w.Write([]byte("asdf"))

	case <-ctx.Done():
		fmt.Printf("cancelled")
		w.Write([]byte("cancelled"))
		return
	}
}
