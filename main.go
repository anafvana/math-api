package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func calculate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "ERROR: Request must be of type POST", http.StatusMethodNotAllowed)
		return
	}

	// Read request body
	byteBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf(string(byteBody))
}

func main() {
	host := "localhost"
	port := "3000"

	fmt.Println(fmt.Sprintf("Opening server on %s:%s", host, port))

	http.HandleFunc("/", calculate)
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil)
	log.Fatal(err)
}
