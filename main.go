package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)

	fmt.Println("Server is running on 5000")

	err := http.ListenAndServe(":5000", mux)

	if err != nil {
		fmt.Println("Server error", err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to go server!")
}
