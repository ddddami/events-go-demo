package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP server port")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	fmt.Printf("Running server on %s\n", *addr)
	http.ListenAndServe(*addr, mux)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hiya")
}
