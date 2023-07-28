package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Handler for the root path '/'
func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", rootHandler)
	fmt.Println("Server is running on port 5000 with metrics on /metrics")
	http.ListenAndServe(":5000", nil)
}
