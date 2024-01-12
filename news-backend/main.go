package main

import (
	"fmt"
	"log"
	"net/http"
)

func routes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "News App")
	})
}

func main() {
	routes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
