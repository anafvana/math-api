package api

import (
	"errors"
	"fmt"
	"reflect"
)

// CheckFields checks that a struct of type T contains non-null values for every type T field
func CheckFields[T any](data T) error {
	datatype := reflect.TypeOf(data)

	// Throw error if attempting to evaluate non-struct
	if datatype.Kind() != reflect.Struct {
		return errors.New("Data is not of kind struct")
	}

	// Check each struct field for null values and build error message
	err := ""
	errCount := 0
	for i := 0; i < datatype.NumField(); i++ {
		fieldName := datatype.Field(i).Name
		if reflect.ValueOf(data).FieldByName(fieldName).IsNil() {
			err += fmt.Sprintf("\nMissing field %s", fieldName)
			errCount++
		}
	}

	// Return nil or error(s)
	if err == "" {
		return nil
	} else {
		return errors.New(fmt.Sprintf("%d ERRORS FOUND: %s", errCount, err))
	}
}
