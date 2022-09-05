package api_test

import (
	"math-api/api"
	"testing"
)

func TestCheckField(t *testing.T) {
	type Test struct {
		hi    *string
		hello *int
		bye   *float64
	}

	hi := "hi"
	hello := 1
	bye := 2.2

	err := api.CheckFields(Test{&hi, &hello, &bye})

	if err != nil {
		t.Logf("Got %s but expected no error", err)
		t.Fail()
	}

	err = api.CheckFields(Test{
		hello: &hello,
		bye:   &bye,
	})

	if err == nil {
		t.Logf("Got no error, but should have")
		t.Fail()
	}

	err = api.CheckFields(hi)

	if err == nil {
		t.Logf("Got no error, but should have")
		t.Fail()
	}
}
