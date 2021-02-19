package web

import (
	"fmt"
	"net/http"
)

// Start a web server
func Start(addr string) {
	routes()
	http.ListenAndServe(addr, nil)
}

func routes() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello, universe")
}
