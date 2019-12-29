package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"net/http"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 7036, "listen port")
	flag.Parse()

	mux := http.NewServeMux()

	port := os.Getenv("PORT")
	if ; port=="" {
		port = "7036"
	}
	addr := fmt.Sprintf(":%s", port)
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, req.RemoteAddr)
	})
	log.Fatal(http.ListenAndServe(addr, mux))
}
