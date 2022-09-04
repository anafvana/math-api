package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type MathRequest struct {
	FirstNumber  *int    `json:"first_number"`
	SecondNumber *int    `json:"second_number"`
	Operation    *string `json:"operation"`
}

type MathResponse struct {
	Result int `json:"result"`
}

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

	dec := json.NewDecoder(bytes.NewReader(byteBody))
	dec.DisallowUnknownFields()
	var body MathRequest
	if err = dec.Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var result int
	switch strings.ToUpper(*body.Operation) {
	case "ADD":
		result = *body.FirstNumber + *body.SecondNumber
	case "SUBTRACT":
		result = *body.FirstNumber - *body.SecondNumber
	case "MULTIPLY":
		result = *body.FirstNumber * *body.SecondNumber
	case "DIVIDE":
		if *body.SecondNumber == 0 {
			http.Error(w, "Cannot divide by 0", http.StatusUnprocessableEntity)
			return
		}
		result = *body.FirstNumber / *body.SecondNumber
	default:
		http.Error(w, "Requested operation is invalid. Operation must be ADD, SUBTRACT, MULTIPLY or DIVIDE", http.StatusUnprocessableEntity)
		return
	}

	response := MathResponse{
		Result: result,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	host := "localhost"
	port := "3000"

	fmt.Println(fmt.Sprintf("Opening server on %s:%s", host, port))

	http.HandleFunc("/", calculate)
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil)
	log.Fatal(err)
}
