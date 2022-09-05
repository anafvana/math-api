package api_test

import (
	"math-api/api"
	"testing"
)

func TestRenameGoToJson(t *testing.T) {
	res := api.RenameGoToJson("AnythingANYNothing")
	exp := "anything_any_nothing"

	if res != exp {
		t.Logf("Got %s but expected %s", res, exp)
		t.Fail()
	}

	res = api.RenameGoToJson("Anything")
	exp = "anything"

	if res != exp {
		t.Logf("Got %s but expected %s", res, exp)
		t.Fail()
	}

	res = api.RenameGoToJson("AnythingA")
	exp = "anything_a"

	if res != exp {
		t.Logf("Got %s but expected %s", res, exp)
		t.Fail()
	}

	res = api.RenameGoToJson("anythingANY")
	exp = "anything_any"

	if res != exp {
		t.Logf("Got %s but expected %s", res, exp)
		t.Fail()
	}

	res = api.RenameGoToJson("anything")
	exp = "anything"

	if res != exp {
		t.Logf("Got %s but expected %s", res, exp)
		t.Fail()
	}

	res = api.RenameGoToJson("anything1")
	exp = "anything1"

	if res != exp {
		t.Logf("Got %s but expected %s", res, exp)
		t.Fail()
	}

	res = api.RenameGoToJson("anything1Anything")
	exp = "anything1_anything"

	if res != exp {
		t.Logf("Got %s but expected %s", res, exp)
		t.Fail()
	}

	res = api.RenameGoToJson("anything1ANYTHING1Anything1")
	exp = "anything1_anything1_anything1"

	if res != exp {
		t.Logf("Got %s but expected %s", res, exp)
		t.Fail()
	}

	res = api.RenameGoToJson("anything")

	if res == "" {
		t.Logf("Got %s, which UNEXPECTEDLY matches empty string", res)
		t.Fail()
	}

}

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
