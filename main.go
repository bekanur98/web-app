package main

import (
	"fmt"
	"net/http"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my greatt site</h1>")
}

func main() {
	http.HandleFunc("/", handleHome)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", nil)
}
