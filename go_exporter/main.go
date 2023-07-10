package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, " <p> Metrics are available at <b>/metrics</b> </p> ")
	})
	fmt.Println("Server is running on port 5000 with metrics on /metrics")
	http.ListenAndServe(":5000", nil)
}
