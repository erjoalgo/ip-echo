package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"github.com/kr/pretty"
)

func ReadUserIP(r *http.Request) string {
	// https://stackoverflow.com/questions/27234861/
	ipAddress := r.Header.Get("X-Real-Ip")
	if ipAddress == "" {
		ipAddress = r.Header.Get("X-Forwarded-For")
	}
	if ipAddress == "" {
		ipAddress = r.RemoteAddr
	}
	return ipAddress
}

func main() {
	var port int
	flag.IntVar(&port, "port", 7036, "listen port")
	flag.Parse()

	mux := http.NewServeMux()
	addr := fmt.Sprintf(":%d", port)
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		ipAdd := ReadUserIP(req)
		fmt.Fprintln(w, ipAdd)
		fmt.Fprintln(w, "")
		fmt.Fprintf(w, "%# v", pretty.Formatter(req))
	})
	fmt.Printf("listening on %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
