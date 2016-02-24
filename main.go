package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		// w.WriteHeader(200)
		fmt.Fprintf(w, req.RemoteAddr)

	})
	log.Fatal(http.ListenAndServe(":7036", mux))
}
