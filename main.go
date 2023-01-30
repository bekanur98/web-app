package main

import (
	"fmt"
	"net/http"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	fmt.Fprint(w, "<h1>Welcome to my greatt site</h1>")
}

func main() {
	http.HandleFunc("/", handleHome)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", nil)
}
