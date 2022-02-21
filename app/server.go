package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	http.Handle("/metrics", promhttp.Handler())
	log.Println("Starting server on port 3000")
	log.Fatalln(http.ListenAndServe(":3000", nil))
}
