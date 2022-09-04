package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	host := "localhost"
	port := "3000"

	fmt.Println(fmt.Sprintf("Opening server on %s:%s", host, port))

	http.HandleFunc("/", handleRequest)
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil)
	log.Fatal(err)
}
