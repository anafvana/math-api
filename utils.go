package main

import (
	"errors"
	"fmt"
	"reflect"
)

func CheckFields[T any](data T) error {
	datatype := reflect.TypeOf(data)

	if datatype.Kind() != reflect.Struct {
		return errors.New("Data is not of kind struct")
	}

	err := ""
	errCount := 0
	for i := 0; i < datatype.NumField(); i++ {
		fieldName := datatype.Field(i).Name
		if reflect.ValueOf(data).FieldByName(fieldName).IsNil() {
			err += fmt.Sprintf("\nMissing field %s", fieldName)
			errCount++
		}
	}

	if err == "" {
		return nil
	} else {
		return errors.New(fmt.Sprintf("%d ERRORS FOUND: %s", errCount, err))
	}
}
