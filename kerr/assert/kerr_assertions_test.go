package assert

import (
	"testing"

	"errors"

	"kego.io/kerr"
)

func TestIsError(t *testing.T) {

	a := New(new(testing.T))

	var e = error(nil)

	if a.IsError(e, "foo", "bar") {
		t.Error("a.IsError should return false")
	}

	if a.HasError(e, "foo", "bar") {
		t.Error("a.HasError should return false")
	}

	e = kerr.New("HKROXGJXIV", "")

	if a.IsError(e, "foo", "bar") {
		t.Error("a.IsError should return false")
	}

	if !a.IsError(e, "HKROXGJXIV", "bar") {
		t.Error("a.IsError should return true")
	}

	e1 := kerr.Wrap("JGAOOWPMNE", e)

	if a.IsError(e1, "HKROXGJXIV", "bar") {
		t.Error("a.IsError should return false")
	}

	if !a.IsError(e1, "JGAOOWPMNE", "bar") {
		t.Error("a.IsError should return true")
	}

	if !a.HasError(e1, "HKROXGJXIV", "bar") {
		t.Error("a.HasError should return true")
	}

	if a.HasError(e1, "foo", "bar") {
		t.Error("a.HasError should return false")
	}

	e = errors.New("foo")

	if a.IsError(e, "foo", "bar") {
		t.Error("a.IsError should return false")
	}

	if a.HasError(e, "foo", "bar") {
		t.Error("a.HasError should return false")
	}

}
