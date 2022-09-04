package main

import (
	"fmt"
	"log"
	"net/http"
)

func calculate(w http.ResponseWriter, r *http.Request) {

}

func main() {
	host := "localhost"
	port := "3000"

	fmt.Println(fmt.Sprintf("Opening server on %s:%s", host, port))

	http.HandleFunc("/", calculate)
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil)
	log.Fatal(err)
}
