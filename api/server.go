package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func calculate(body MathRequest) (int, error) {
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
			return result, errors.New("Cannot divide by 0")
		}
		result = *body.FirstNumber / *body.SecondNumber
	default:
		return result, errors.New("Requested operation is invalid. Operation must be ADD, SUBTRACT, MULTIPLY or DIVIDE")
	}
	return result, nil
}

// HandleRequest handles the incoming request, validates method, fields and data, and returns the operation's result (or an error)
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	placeholder := `{
        "first_number": integer,
        "second_number": integer,
        "operation": string
    }`

	// Disallow verbs other than POST
	if r.Method != "POST" {
		http.Error(w, "ERROR: Request must be of type POST", http.StatusMethodNotAllowed)
		return
	}

	// Read request body
	byteBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf(`%s
Expected: 
    %s`, err.Error(), placeholder), http.StatusBadRequest)
		return
	}

	// Decode body
	dec := json.NewDecoder(bytes.NewReader(byteBody))
	dec.DisallowUnknownFields()
	var body MathRequest
	if err = dec.Decode(&body); err != nil {
		http.Error(w, fmt.Sprintf(`%s
Expected: 
    %s`, err.Error(), placeholder), http.StatusBadRequest)
		return
	}

	// Ensure all fields are non-nil
	if err = CheckFields(body); err != nil {
		http.Error(w, fmt.Sprintf(`%s
Expected: 
    %s`, err.Error(), placeholder), http.StatusBadRequest)
		return
	}

	// Run desired mathematical operation
	result, err := calculate(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	// Generate and send response
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
