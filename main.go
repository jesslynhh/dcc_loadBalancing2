package main

import (
	"fmt"
	"log"
	"net/http"
)

func vm1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "VM 1 Page 2\n")
	fmt.Fprint(w, "Next endpoint: /yay")
}

func vm2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "yay")
}

func handleRequests() {
	log.Print("Server Running at: http://localhost:8080")
	http.HandleFunc("/", vm1)
	http.HandleFunc("/yay", vm2)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequests()
}
