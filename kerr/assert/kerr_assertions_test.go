package assert

import (
	"testing"

	"errors"

	"kego.io/kerr"
)

func TestIsError(t *testing.T) {

	assert := New(new(testing.T))

	var e = error(nil)

	if assert.IsError(e, "foo", "bar") {
		t.Error("assert.IsError should return false")
	}

	if assert.HasError(e, "foo", "bar") {
		t.Error("assert.HasError should return false")
	}

	e = kerr.New("HKROXGJXIV", nil, "", "")

	if assert.IsError(e, "foo", "bar") {
		t.Error("assert.IsError should return false")
	}

	if !assert.IsError(e, "HKROXGJXIV", "bar") {
		t.Error("assert.IsError should return true")
	}

	e1 := kerr.New("JGAOOWPMNE", e, "baz", "qux")

	if assert.IsError(e1, "HKROXGJXIV", "bar") {
		t.Error("assert.IsError should return false")
	}

	if !assert.IsError(e1, "JGAOOWPMNE", "bar") {
		t.Error("assert.IsError should return true")
	}

	if !assert.HasError(e1, "HKROXGJXIV", "bar") {
		t.Error("assert.HasError should return true")
	}

	if assert.HasError(e1, "foo", "bar") {
		t.Error("assert.HasError should return false")
	}

	e = errors.New("foo")

	if assert.IsError(e, "foo", "bar") {
		t.Error("assert.IsError should return false")
	}

	if assert.HasError(e, "foo", "bar") {
		t.Error("assert.HasError should return false")
	}

}
