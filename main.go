package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 7036, "listen port")
	flag.Parse()

	mux := http.NewServeMux()
	addr := fmt.Sprintf(":%d", port)
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, req.RemoteAddr)
	})
	fmt.Printf("listening on %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
