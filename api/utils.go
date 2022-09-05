package api

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/dlclark/regexp2"
)

// RenameGoToJson transforms a Go-style name (HelloWorld or helloWorld) into a JSON-style name (hello_world)
// It assumes variables will contain only 0-9, a-z and A-Z
// All-caps words should be handled correctly, as long as they are followed by a PascalCase word or are at the end of a sentence (helloWORLD or HELLOWorld = hello-world)
func RenameGoToJson(golang string) string {
	// Compile regex
	re := regexp2.MustCompile(`((?<=[a-z0-9])[A-Z]|[A-Z](?=[a-z]))`, regexp2.None) //regexp2 required for look-behind and look-ahead
	match, _ := re.FindStringMatch(golang)

	// Implement regexp.Split (unavailable under regexp2)
	prev := 0
	last := len(golang)
	separator := "_"
	json := ""

	for match != nil {
		json += strings.ToLower(golang[prev:match.Index])
		prev = match.Index
		match, _ = re.FindNextMatch(match)

		if match != nil && prev != 0 && match.Index != last {
			json += separator
		}
	}

	// Ensure last word (or only word) is added
	if prev != 0 {
		json += separator
	}
	json += strings.ToLower(golang[prev:last])

	return json
}

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
			err += fmt.Sprintf("\nMissing field %s", RenameGoToJson(fieldName)) // Field name converted to JSON standard
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
