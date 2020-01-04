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
	var verboseReplies bool
	flag.IntVar(&port, "port", 7036, "listen port")
	flag.BoolVar(&verboseReplies, "verbose-replies", false,
		"whether to serve replies with excessive debug information")
	flag.Parse()

	mux := http.NewServeMux()
	addr := fmt.Sprintf(":%d", port)
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		ipAdd := ReadUserIP(req)
		fmt.Fprintf(w, "Ip-Address: %s\n", ipAdd)
		fmt.Fprintf(w, "User-Agent: %s\n", req.Header.Get("User-Agent"))
		if verboseReplies {
			fmt.Fprintln(w, "")
			fmt.Fprintf(w, "%# v", pretty.Formatter(req))
		}
	})
	fmt.Printf("listening on %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
