package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"github.com/kr/pretty"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 7036, "listen port")
	flag.Parse()

	mux := http.NewServeMux()
	addr := fmt.Sprintf(":%d", port)
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, req.RemoteAddr)
		fmt.Fprintln(w, "")
		fmt.Fprintf(w, "%# v", pretty.Formatter(req))
	})
	fmt.Printf("listening on %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
